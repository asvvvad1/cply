# cply
Search and copy lyrics from your terminal
This version is a rewrite of my [php script](https://github.com/asvvvad/cply-php) in Go

## Requirements
I only tested it on a ubuntu based distro but should work on all others including BSD and MacOS and Windows

Linux users need xclip or xsel to be installed for the copying to work

## Install


First, generate an api token for the search functionality (required): https://genius.com/api-clients

```bash
git clone https://github.com/asvvvad/cply
cd cply
```
Then edit [main.go](main.go#L74) with the generated token
```bash
nano +74 main.go
```
Finally install with the Go toolchain:
`go install`

## Usage:
[![asciicast](https://asciinema.org/a/320915.svg)](https://asciinema.org/a/320915)


### Modules used:
- [tsirysndr/go-genius](github.com/tsirysndr/go-genius)
- [bigheadgeorge/lyrics/genius](github.com/bigheadgeorge/lyrics/genius)
- [paulrademacher/climenu](github.com/paulrademacher/climenu)
- [wzshiming/ctc](github.com/wzshiming/ctc)
- [atotto/clipboard](github.com/atotto/clipboard)
