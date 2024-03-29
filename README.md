<p align="center"><a href="#readme"><img src="https://gh.kaos.st/redis-monitor-top.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/w/redis-monitor-top/ci"><img src="https://kaos.sh/w/redis-monitor-top/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/r/redis-monitor-top"><img src="https://kaos.sh/r/redis-monitor-top.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/b/redis-monitor-top"><img src="https://kaos.sh/b/98c9f6ab-999c-498c-980f-44859b18aae7.svg" alt="codebeat badge" /></a>
  <a href="https://kaos.sh/w/redis-monitor-top/codeql"><img src="https://kaos.sh/w/redis-monitor-top/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#usage-demo">Usage demo</a> • <a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

Tiny Redis client for aggregating stats from MONITOR flow.

### Usage demo

[![demo](https://gh.kaos.st/redis-monitor-top-100.gif)](#usage-demo)

### Installation

#### From source

To build the `redis-monitor-top` from scratch, make sure you have a working Go 1.17+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get github.com/essentialkaos/redis-monitor-top
```

If you want to update `redis-monitor-top` to latest stable release, do:

```
go get -u github.com/essentialkaos/redis-monitor-top
```

#### From [ESSENTIAL KAOS Public Repository](https://yum.kaos.st)

```bash
sudo yum install -y https://yum.kaos.st/get/$(uname -r).rpm
sudo yum install redis-monitor-top
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux from [EK Apps Repository](https://apps.kaos.st/redis-monitor-top/latest).

To install the latest prebuilt version, do:

```bash
bash <(curl -fsSL https://apps.kaos.st/get) redis-monitor-top
```

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
  Start monitoring instance on 192.168.0.123:6821 with 30 second interval and renamed MONITOR command

```

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://kaos.sh/w/redis-monitor-top/ci.svg?branch=master)](https://kaos.sh/w/redis-monitor-top/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/redis-monitor-top/ci.svg?branch=master)](https://kaos.sh/w/redis-monitor-top/ci?query=branch:develop) |

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
