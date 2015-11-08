# Tsunami
An advanced HTTP flooder written in Golang

__This project is huge WIP__

## Basic Usage
```bash
./tsunami http://example.com --threads 100
```

## Help
```bash
./tsunami --help
```

## Quick install
```bash
git clone https://github.com/ammario/tsunami
cd tsunami
export GOPATH=`pwd`
go get ./...
go build
```

## Features
- Live attack stats
- Customizable mutlithreading
- HTTPS support __(Note: Certificates aren't verified for performance)__

## Todo
 - Dynamic tokens (E.g {RANDOM_STRING}, {RANDOM_INT})
 - User Agent randomization
 - All HTTP methods
 - Custom headers
 - Pretty display of attack stats
 - ???