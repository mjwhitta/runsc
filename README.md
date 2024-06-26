# runsc

[![Yum](https://img.shields.io/badge/-Buy%20me%20a%20cookie-blue?labelColor=grey&logo=cookiecutter&style=for-the-badge)](https://www.buymeacoffee.com/mjwhitta)

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/runsc?style=for-the-badge)](https://goreportcard.com/report/github.com/mjwhitta/runsc)
![License](https://img.shields.io/github/license/mjwhitta/runsc?style=for-the-badge)

## What is this?

This Go module allows you to inject shellcode using the Windows Native
API. At this time, there aren't a whole lot of options. In the future,
it will be more configurable via an a la carte style launcher.

## How to install

Open a terminal and run the following:

```
$ git clone https://github.com/mjwhitta/runsc.git
$ cd ./runsc
$ git submodule update --init
$ make GOOS=windows cgo
```

## Usage

See [main.go](cmd/runsc/main.go) for example usage.

## Links

- [Source](https://github.com/mjwhitta/runsc)
