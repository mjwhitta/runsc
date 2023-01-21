# runsc

<a href="https://www.buymeacoffee.com/mjwhitta">🍪 Buy me a cookie</a>

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/runsc)](https://goreportcard.com/report/github.com/mjwhitta/runsc)

## What is this?

This Go module allows you to inject shellcode using the Windows Native
API (specifically functions that are not hooked by EDR solutions).

## How to install

Open a terminal and run the following:

```
$ go get --ldflags="-s -w" --trimpath -u github.com/mjwhitta/runsc
```

## Usage

See [main.go](cmd/runsc/main.go) for example usage.

## Links

- [Source](https://github.com/mjwhitta/runsc)
