## Redis Monitor Top [![Build Status](https://travis-ci.org/essentialkaos/redis-monitor-top.svg?branch=master)](https://travis-ci.org/essentialkaos/redis-monitor-top) [![Go Report Card](https://goreportcard.com/badge/github.com/essentialkaos/redis-monitor-top)](https://goreportcard.com/report/github.com/essentialkaos/redis-monitor-top) [![License](https://gh.kaos.io/ekol.svg)](https://essentialkaos.com/ekol)

Tiny Redis client for aggregating stats from MONITOR flow.

### Installation

#### From source

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)):

```
git config --global http.https://pkg.re.followRedirects true
```

To build the `redis-monitor-top` from scratch, make sure you have a working Go 1.6+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get github.com/essentialkaos/redis-monitor-top
```

If you want to update `redis-monitor-top` to latest stable release, do:

```
go get -u github.com/essentialkaos/redis-monitor-top
```

#### From ESSENTIAL KAOS Public repo for RHEL6/CentOS6

```bash
[sudo] yum install -y https://yum.kaos.io/6/release/x86_64/kaos-repo-8.0-0.el6.noarch.rpm
[sudo] yum install redis-monitor-top
```

#### From ESSENTIAL KAOS Public repo for RHEL7/CentOS7

```bash
[sudo] yum install -y https://yum.kaos.io/7/release/x86_64/kaos-repo-8.0-0.el7.noarch.rpm
[sudo] yum install redis-monitor-top
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and OS X from [EK Apps Repository](https://apps.kaos.io/redis-monitor-top/latest).

### Usage

```
Usage: redis-monitor-top {options} command

Options

  --host, -h ip/host         Server hostname (127.0.0.1 by default)
  --port, -p port            Server port (6379 by default)
  --password, -a password    Password to use when connecting to the server
  --timeout, -t 1-300        Connection timeout in seconds (3 by default)
  --interval, -i 1-3600      Interval in seconds (60 by default)
  --no-color, -nc            Disable colors in output
  --help                     Show this help message
  --version, -v              Show version

Examples

  redis-monitor-top -h 192.168.0.123 -p 6821 -t 15 MONITOR
  Start monitoring instance on 192.168.0.123:6821 with 15 second timeout

  redis-monitor-top -h 192.168.0.123 -p 6821 -i 30 MY_MONITOR
  Start monitoring instance on 192.168.0.123:6821 with 30 second interval with renamed MONITOR command

```

### Build Status

| Repository | Status |
|------------|--------|
| Stable | [![Build Status](https://travis-ci.org/essentialkaos/redis-monitor-top.svg?branch=master)](https://travis-ci.org/essentialkaos/redis-monitor-top) |
| Unstable | [![Build Status](https://travis-ci.org/essentialkaos/redis-monitor-top.svg?branch=develop)](https://travis-ci.org/essentialkaos/redis-monitor-top) |

### License

[EKOL](https://essentialkaos.com/ekol)
