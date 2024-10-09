<p align="center"><a href="#readme"><img src=".github/images/card.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/w/redis-monitor-top/ci"><img src="https://kaos.sh/w/redis-monitor-top/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/r/redis-monitor-top"><img src="https://kaos.sh/r/redis-monitor-top.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/w/redis-monitor-top/codeql"><img src="https://kaos.sh/w/redis-monitor-top/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src=".github/images/license.svg"/></a>
</p>

<p align="center"><a href="#usage-demo">Usage demo</a> • <a href="#installation">Installation</a> • <a href="#usage">Usage</a> • <a href="#ci-status">CI Status</a> • <a href="#license">License</a></p>

<br/>

Tiny Redis client for aggregating stats from MONITOR flow.

### Usage demo

[![demo](https://gh.kaos.st/redis-monitor-top-100.gif)](#usage-demo)

### Installation

#### From source

To build the `redis-monitor-top` from scratch, make sure you have a working [Go 1.22+](https://github.com/essentialkaos/.github/blob/master/GO-VERSION-SUPPORT.md) workspace (_[instructions](https://go.dev/doc/install)_), then:

```
go install github.com/essentialkaos/redis-monitor-top@latest
```

#### From [ESSENTIAL KAOS Public Repository](https://kaos.sh/kaos-repo)

```bash
sudo dnf install -y https://pkgs.kaos.st/kaos-repo-latest.el$(grep 'CPE_NAME' /etc/os-release | tr -d '"' | cut -d':' -f5).noarch.rpm
sudo dnf install redis-monitor-top
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux from [EK Apps Repository](https://apps.kaos.st/redis-monitor-top/latest).

To install the latest prebuilt version, do:

```bash
bash <(curl -fsSL https://apps.kaos.st/get) redis-monitor-top
```

### Usage

<img src=".github/images/usage.svg" />

### CI Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://kaos.sh/w/redis-monitor-top/ci.svg?branch=master)](https://kaos.sh/w/redis-monitor-top/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/redis-monitor-top/ci.svg?branch=master)](https://kaos.sh/w/redis-monitor-top/ci?query=branch:develop) |

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
