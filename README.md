# runsc

[![Yum](https://img.shields.io/badge/-Buy%20me%20a%20cookie-blue?labelColor=grey&logo=cookiecutter&style=for-the-badge)](https://www.buymeacoffee.com/mjwhitta)

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/runsc?style=for-the-badge)](https://goreportcard.com/report/github.com/mjwhitta/runsc)
![License](https://img.shields.io/github/license/mjwhitta/runsc?style=for-the-badge)

## What is this?

This Go module allows you to inject shellcode using the Windows Native
API (specifically functions that are not hooked by EDR solutions).

## How to install

Open a terminal and run the following:

```
$ go get --ldflags "-s -w" --trimpath -u github.com/mjwhitta/runsc
```

## Usage

See [main.go](cmd/runsc/main.go) for example usage.

## Links

- [Source](https://github.com/mjwhitta/runsc)
