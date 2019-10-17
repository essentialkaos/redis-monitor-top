// +build linux

package main

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"

	"pkg.re/essentialkaos/ek.v11/fmtc"
	"pkg.re/essentialkaos/ek.v11/fmtutil"
	"pkg.re/essentialkaos/ek.v11/fmtutil/table"
	"pkg.re/essentialkaos/ek.v11/mathutil"
	"pkg.re/essentialkaos/ek.v11/options"
	"pkg.re/essentialkaos/ek.v11/strutil"
	"pkg.re/essentialkaos/ek.v11/system/procname"
	"pkg.re/essentialkaos/ek.v11/timeutil"
	"pkg.re/essentialkaos/ek.v11/usage"
	"pkg.re/essentialkaos/ek.v11/usage/completion/bash"
	"pkg.re/essentialkaos/ek.v11/usage/completion/fish"
	"pkg.re/essentialkaos/ek.v11/usage/completion/zsh"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	APP  = "Redis Monitor Top"
	VER  = "1.3.1"
	DESC = "Tiny Redis client for aggregating stats from MONITOR flow"
)

const (
	OPT_HOST     = "h:host"
	OPT_PORT     = "p:port"
	OPT_AUTH     = "a:password"
	OPT_TIMEOUT  = "t:timeout"
	OPT_INTERVAL = "i:interval"
	OPT_NO_COLOR = "nc:no-color"
	OPT_HELP     = "help"
	OPT_VER      = "v:version"

	OPT_COMPLETION = "completion"
)

const MAX_COMMANDS = 128

// ////////////////////////////////////////////////////////////////////////////////// //

// CommandInfo contains name of command and count
type CommandInfo struct {
	Name  string
	Count int64
}

type CommandInfoSlice []*CommandInfo

func (s CommandInfoSlice) Len() int           { return len(s) }
func (s CommandInfoSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s CommandInfoSlice) Less(i, j int) bool { return s[i].Count < s[j].Count }

type Stats struct {
	Data  map[string]*CommandInfo
	Slice CommandInfoSlice

	Dirty   bool
	HasData bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

// optMap is map with options
var optMap = options.Map{
	OPT_HOST:     {Type: options.MIXED, Value: "127.0.0.1"},
	OPT_PORT:     {Value: "6379"},
	OPT_TIMEOUT:  {Type: options.INT, Value: 3, Min: 1, Max: 300},
	OPT_AUTH:     {},
	OPT_INTERVAL: {Type: options.INT, Value: 60, Min: 1, Max: 3600},
	OPT_NO_COLOR: {Type: options.BOOL},
	OPT_HELP:     {Type: options.BOOL, Alias: "u:usage"},
	OPT_VER:      {Type: options.BOOL, Alias: "ver"},

	OPT_COMPLETION: {},
}

var conn net.Conn
var stats *Stats

// ////////////////////////////////////////////////////////////////////////////////// //

// main is main function
func main() {
	runtime.GOMAXPROCS(4)

	args, errs := options.Parse(optMap)

	if len(errs) != 0 {
		for _, err := range errs {
			printError(err.Error())
		}

		os.Exit(1)
	}

	if options.Has(OPT_COMPLETION) {
		genCompletion()
	}

	if options.GetB(OPT_NO_COLOR) {
		fmtc.DisableColors = true
	}

	if options.GetB(OPT_VER) {
		showAbout()
		return
	}

	if options.GetB(OPT_HELP) || options.GetS(OPT_HOST) == "true" {
		showUsage()
		return
	}

	cmd := "MONITOR"

	if len(args) != 0 && strings.ToUpper(args[0]) != cmd {
		cmd = strutil.Copy(args[0])
		maskCommand(args[0])
	}

	start(cmd)
}

// maskCommand mask command in process tree
func maskCommand(cmd string) {
	cmdLen := mathutil.Max(len(cmd), 16)
	procname.Replace(cmd, strings.Repeat("*", cmdLen))
}

// start connect to redis and starts monitor flow processing
func start(cmd string) {
	err := connectToRedis(
		options.GetS(OPT_HOST),
		options.GetS(OPT_PORT),
		time.Second*time.Duration(options.GetI(OPT_TIMEOUT)),
	)

	if err != nil {
		printErrorAndExit(err.Error())
	}

	processCommands(cmd)
}

// connectToRedis connect to redis instance
func connectToRedis(host, port string, timeout time.Duration) error {
	var err error

	conn, err = net.DialTimeout("tcp", host+":"+port, timeout)

	if err != nil {
		return err
	}

	if options.GetS(OPT_AUTH) != "" {
		_, err = conn.Write([]byte("AUTH " + options.GetS(OPT_AUTH) + "\n"))

		if err != nil {
			return err
		}
	}

	return nil
}

// processCommands send monitor command to redis and process command flow
func processCommands(cmd string) {
	connbuf := bufio.NewReader(conn)
	conn.Write([]byte(cmd + "\n"))

	stats = NewStats()

	go printStats()

	for {
		str, err := connbuf.ReadString('\n')

		if len(str) > 0 {
			if str == "+OK\r\n" {
				continue
			}

			if strings.HasPrefix(str, "-ERR ") {
				printErrorAndExit("Redis return error message: " + strings.TrimRight(str[1:], "\r\n"))
			}

			if stats.Dirty {
				stats.Clean()
			}

			stats.Increment(extractCommandName(str))
		}

		if err != nil {
			printErrorAndExit(err.Error())
		}
	}
}

// printStats periodically print stats
func printStats() {
	last := time.Now()
	interval := time.Second * time.Duration(options.GetI(OPT_INTERVAL))
	t := table.NewTable("DATE & TIME", "COUNT", "RPS", "COMMAND")
	t.SetSizes(20, 10, 10)
	t.SetAlignments(table.ALIGN_RIGHT, table.ALIGN_RIGHT, table.ALIGN_RIGHT)

	for range time.NewTicker(time.Millisecond * 250).C {
		if time.Since(last) >= interval {
			renderStats(t)
			last = time.Now()
		}
	}
}

// renderStats calculate and render stats
func renderStats(t *table.Table) {
	now := time.Now()

	if !stats.HasData || stats.Dirty {
		t.Print(
			timeutil.Format(now, "%Y/%m/%d %H:%M:%S"),
			"{s-}----------{!}",
			"{s-}----------{!}",
			"{s-}----------{!}",
		)
		t.Separator()
		return
	}

	sort.Sort(sort.Reverse(stats.Slice))

	interval := float64(options.GetI(OPT_INTERVAL))

	for i, info := range stats.Slice {
		if info.Count == 0 {
			break
		}

		if i == 0 {
			t.Print(
				timeutil.Format(now, "%Y/%m/%d %H:%M:%S"),
				fmtutil.PrettyNum(info.Count),
				fmtutil.PrettyNum(formatFloat(float64(info.Count)/interval)),
				strings.ToUpper(info.Name),
			)
		} else {
			t.Print(
				" ",
				fmtutil.PrettyNum(info.Count),
				fmtutil.PrettyNum(formatFloat(float64(info.Count)/interval)),
				strings.ToUpper(info.Name),
			)
		}

	}

	t.Separator()

	stats.Dirty = true
}

// extractCommandName extract command name from full command
func extractCommandName(command string) string {
	cmdStart := strings.IndexRune(command, ']')

	if cmdStart == -1 {
		return ""
	}

	cmdStart += 3

	cmdEnd := strings.IndexRune(command[cmdStart:], '"')

	if cmdEnd == -1 {
		return ""
	}

	return command[cmdStart : cmdStart+cmdEnd]
}

// formatFloat format floating numbers
func formatFloat(f float64) float64 {
	switch {
	case f > 500:
		return mathutil.Round(f, 0)
	case f > 50:
		return mathutil.Round(f, 1)
	case f > 0.3:
		return mathutil.Round(f, 2)
	}

	return f
}

// printErrorAndExit print error message and exit from utility
func printErrorAndExit(f string, a ...interface{}) {
	printError(f, a...)
	shutdown(1)
}

// printError prints error message to console
func printError(f string, a ...interface{}) {
	fmtc.Fprintf(os.Stderr, "{r}"+f+"{!}\n", a...)
}

// shutdown close connection to Redis and exit from utility
func shutdown(code int) {
	if conn != nil {
		conn.Close()
	}

	os.Exit(code)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewStats create new stats struct
func NewStats() *Stats {
	return &Stats{
		Data:  make(map[string]*CommandInfo),
		Slice: make([]*CommandInfo, 0),
	}
}

// Clean clean stats
func (s *Stats) Clean() {
	if !s.Dirty {
		return
	}

	for _, info := range s.Data {
		info.Count = 0
	}

	s.HasData = false
	s.Dirty = false
}

// Increment increment counter for given command
func (s *Stats) Increment(command string) {
	if s.Data[command] == nil {
		info := &CommandInfo{command, 0}
		s.Data[command] = info
		s.Slice = append(s.Slice, info)
	}

	s.Data[command].Count++
	s.HasData = true
}

// ////////////////////////////////////////////////////////////////////////////////// //

// showUsage print usage info
func showUsage() {
	genUsage().Render()
}

// genUsage generates usage info
func genUsage() *usage.Info {
	info := usage.NewInfo("", "command")

	info.AddOption(OPT_HOST, "Server hostname {s-}(127.0.0.1 by default){!}", "ip/host")
	info.AddOption(OPT_PORT, "Server port {s-}(6379 by default){!}", "port")
	info.AddOption(OPT_AUTH, "Password to use when connecting to the server", "password")
	info.AddOption(OPT_TIMEOUT, "Connection timeout in seconds {s-}(3 by default){!}", "1-300")
	info.AddOption(OPT_INTERVAL, "Interval in seconds {s-}(60 by default){!}", "1-3600")
	info.AddOption(OPT_NO_COLOR, "Disable colors in output")
	info.AddOption(OPT_HELP, "Show this help message")
	info.AddOption(OPT_VER, "Show version")

	info.AddExample(
		"-h 192.168.0.123 -p 6821 -t 15 MONITOR",
		"Start monitoring instance on 192.168.0.123:6821 with 15 second timeout",
	)

	info.AddExample(
		"-h 192.168.0.123 -p 6821 -i 30 MY_MONITOR",
		"Start monitoring instance on 192.168.0.123:6821 with 30 second interval and renamed MONITOR command",
	)

	return info
}

// genCompletion generates completion for different shells
func genCompletion() {
	info := genUsage()

	switch options.GetS(OPT_COMPLETION) {
	case "bash":
		fmt.Printf(bash.Generate(info, "redis-monitor-top"))
	case "fish":
		fmt.Printf(fish.Generate(info, "redis-monitor-top"))
	case "zsh":
		fmt.Printf(zsh.Generate(info, optMap, "redis-monitor-top"))
	default:
		os.Exit(1)
	}

	os.Exit(0)
}

// showAbout print info about version
func showAbout() {
	about := &usage.About{
		App:     APP,
		Version: VER,
		Desc:    DESC,
		Year:    2006,
		Owner:   "ESSENTIAL KAOS",
		License: "Essential Kaos Open Source License <https://essentialkaos.com/ekol>",
	}

	about.Render()
}
