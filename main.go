package main

import (
	"fmt"
	"github.com/tsirysndr/go-genius"
	lyrics "github.com/bigheadgeorge/lyrics/genius"
	"github.com/paulrademacher/climenu"
	"github.com/wzshiming/ctc"
	"github.com/atotto/clipboard"
	"os"
	"strings"
	"strconv"
	"regexp"
)

func err(error string, exit bool) {
	fmt.Println(ctc.ForegroundBrightRed, "Error: ", ctc.Reset, error)
	if exit {
		os.Exit(1)
	}
}

func main() {
	help := fmt.Sprint(ctc.ForegroundBrightCyan, "CPLY:", ctc.Reset, " Get lyrics copied to your clipboard.\n\n", ctc.ForegroundBrightGreen|ctc.Underline, "Usage:", ctc.Reset, " cply [options] song title..\n\n", ctc.ForegroundBrightGreen|ctc.Underline, "Options:", ctc.Reset, "\n\n    ", ctc.ForegroundBrightGreen, "-h, --help:", ctc.Reset, "     Print this help text and exit.\n    ", ctc.ForegroundBrightGreen, "-1, --first:", ctc.Reset, "    Do not prompt to choose a song but select the first result.\n    ", ctc.ForegroundBrightGreen, "-p, --print:", ctc.Reset, "    Print the lyrics not just copy them.\n    ", ctc.ForegroundBrightGreen, "-n, --no-color:", ctc.Reset, "    Don't use colors when printing lyrics.","\n\n", ctc.ForegroundBrightGreen|ctc.Underline, "Environment Variables:", ctc.Reset, "\n\n    ",ctc.ForegroundBrightGreen, "CPLY_FIRST", ctc.Reset,ctc.ForegroundBrightGreen, "CPLY_PRINT", ctc.Reset,ctc.ForegroundBrightGreen, "CPLY_NOCOLOR", ctc.Reset,"\n\nCPLY by ASVVVAD (", ctc.ForegroundBrightCyan, "https://asvvvad.eu.org", ctc.Reset, ") | Github: ", ctc.ForegroundBrightCyan, "(https://github.com/asvvvad/cply)", ctc.Reset)


	if len(os.Args) == 1 {
		err("Song not specified specified\n", false)
		fmt.Println(help)
		os.Exit(1)
	}

	chooseFirst := false
	printLyrics := false
	noColors := false
	q := ""
	// parse flags	

	for i := range os.Args {
		switch os.Args[i] {
			case "-1", "--first":
				chooseFirst = true
			case "-h", "--help":
				fmt.Println(help)
				os.Exit(0)
			case "-p", "--print":
				printLyrics = true
			case "-n", "--no-color":
				noColors = true
		}
	}

	r, _ := regexp.Compile("(-p|--print|-h|--help|-1|--first|./|cply|-n|--no-color)")

	q = r.ReplaceAllString(strings.Join(os.Args, ""), "")

	if q == "" {
		err("Song not specified specified\n", false)
		fmt.Println(help)
		os.Exit(1)
	}

	// check environment variables
	if os.Getenv("CPLY_PRINT") != "" {
		printLyrics = true
	}
	if os.Getenv("CPLY_FIRST") != "" {
		chooseFirst = true
	}
	if os.Getenv("CPLY_NOCOLOR") != "" {
		chooseFirst = true
	}

	token := "access_token"
	client := genius.NewClient(token)
	res, _ := client.Search.Get(q)

	if len(res) == 0 {
		fmt.Println(ctc.ForegroundRed, "No results found", ctc.Reset)
		os.Exit(1)
	}

	url := ""
	title := ""

	if chooseFirst {
		url = res[0].Result.URL
		title = res[0].Result.FullTitle
	} else {
		menu := climenu.NewButtonMenu("Results for "+q, "Choose a song")

		for i := 0; i < len(res); i++  {
			menu.AddMenuItem(res[i].Result.FullTitle, strconv.Itoa(i))
		}

		action, escaped := menu.Run()
		if escaped {
			os.Exit(0)
		}
		actionInt, _ := strconv.ParseInt(action, 0, 64)
		url = res[actionInt].Result.URL
		title = res[actionInt].Result.FullTitle
	}

	
	l, _ := lyrics.LyricsURL(url)
	lyrics := strings.ReplaceAll(fmt.Sprint(title,"\n",strings.Join(l, "\n")), "[", "\n[")

	if clipboard.Unsupported {
		err("can't copy to the clipboard. Here are the lyrics:", false)
		if !noColors {
			lyrics = strings.ReplaceAll(lyrics, "[", fmt.Sprint(ctc.ForegroundBrightYellow,"["))
			lyrics = strings.ReplaceAll(lyrics, "]", fmt.Sprint("]",ctc.Reset))
		}
		fmt.Println(lyrics)
		os.Exit(1)
	} else {
		clipboard.WriteAll(lyrics)
		if (printLyrics) {
			if !noColors {
				lyrics = strings.ReplaceAll(lyrics, "[", fmt.Sprint(ctc.ForegroundBrightYellow,"["))
				lyrics = strings.ReplaceAll(lyrics, "]", fmt.Sprint("]",ctc.Reset))
			}
			fmt.Println(lyrics)
		} else {
			fmt.Println(ctc.ForegroundGreen, "Lyrics written to clipboard successfully", ctc.Reset)
		}
		os.Exit(0)
	}

}