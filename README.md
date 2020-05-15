# cply
Search and copy lyrics from your terminal
This version is a rewrite of my [PHP script](https://github.com/asvvvad/cply-php) in Go

## Requirements/Platforms

I only tested it on a ubuntu based distro but it work on all others including BSD and MacOS and Windows

> Linux users need `xclip` or `xsel` to be installed for the copying to work

> Wayland users need [wl-clipboard](https://github.com/bugaevc/wl-clipboard)

> It _should_ work on Termux but not tested.

## Install

> If you use a Linux based OS with amd64 architecture you can download a pre-build and min-sized binary from [releases](https://github.com/asvvvad/cply/releases/)

First, generate an api token for the search functionality (required): https://genius.com/api-clients

Then, set it in the environment variable `$CPLY_TOKEN` which you should keep in your `~/.profile`

To do that add:  `export CPLY_TOKEN=access_token_here` to end of that file

> changes to that file will only happen when you log out and log in againbut you can run `source ~/.profile` in the shell to test it

Finally just run this command
```bash
go install github.com/asvvvad/cply
```

## Usage:
  - `cply song name and/or artist` search for "song name and/or artist" and gives you results to select from (max. 10)
  - - To select a song simply type its number in the input and press enter, to choose the first one press enter directly or:

  - `cply -first|-1 song name and/or artist` search for "song name and/or artist" then fetch and copy the first result directly
  - - This can be made default by setting the `$CPLY_FIRST` variable

  - `cply -print|-p song name and/or artist` search for "song name and/or artist" print the lyrics instead of just copying
  - - This can be made default by setting the `$CPLY_PRINT` variable
  - - `cply -print|-p -no-color|-n song name and/or artist` print without highlighting (Making the [Chorus] ect yellow)
  - - - This can be made default by setting the `$CPLY_NOCOLOR` variable

  - `cply -1 -p song name and/or artist` search, select first result (if there is), copy it and print.
 
### ASCIICAST
[![asciicast](https://asciinema.org/a/321229.svg)](https://asciinema.org/a/321229)

### Modules used:
- [tsirysndr/go-genius](https://github.com/tsirysndr/go-genius)
- [bigheadgeorge/lyrics/genius](https://github.com/bigheadgeorge/lyrics)
- [wzshiming/ctc](https://github.com/wzshiming/ctc)
- [atotto/clipboard](https://github.com/atotto/clipboard)
