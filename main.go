package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	lyrics "github.com/bigheadgeorge/lyrics/genius"
	"github.com/tsirysndr/go-genius"
	"github.com/wzshiming/ctc"
	"os"
	"strings"
	"flag"
)

func err(error string, exit bool) {
	fmt.Println(ctc.ForegroundBrightRed, "Error: ", ctc.Reset, error)
	if exit {
		os.Exit(1)
	}
}

func info(info string) {
	fmt.Println(ctc.ForegroundBrightYellow, "+ "+info+"\n", ctc.Reset)
}

func success(text string) {
	fmt.Println(ctc.ForegroundGreen, "* "+text+"\n", ctc.Reset)
}

func main() {

	// Get the access_token from the environment to ease the install process
	// suggested by https://www.reddit.com/user/_noobAtLife_/ at https://www.reddit.com/r/golang/comments/g2vh9j/search_and_copy_lyrics_from_the_terminal_first_go/fno9i8k
	if os.Getenv("CPLY_TOKEN") == "" {
		err("Please provide an access token (from https://genius.com/api-clients) in your environment as $CPLY_TOKEN", true)
	}

	help := fmt.Sprint(ctc.ForegroundBrightCyan, "CPLY:", ctc.Reset, " Get lyrics copied to your clipboard.\n\n", ctc.ForegroundBrightGreen|ctc.Underline, "Usage:", ctc.Reset, " cply [options] song title..\n\n", ctc.ForegroundBrightGreen|ctc.Underline, "Options:", ctc.Reset, "\n\n    ", ctc.ForegroundBrightCyan, "-h, -help:", ctc.Reset, "     Print this help text and exit.\n    ", ctc.ForegroundBrightCyan, "-1, -first:", ctc.Reset, "    Do not prompt to choose a song but select the first result.\n    ", ctc.ForegroundBrightCyan, "-p, -print:", ctc.Reset, "    Print the lyrics not just copy them.\n    ", ctc.ForegroundBrightCyan, "-n, -no-color:", ctc.Reset, "    Don't use colors when printing lyrics.", "\n\n", ctc.ForegroundBrightGreen|ctc.Underline, "Environment Variables:", ctc.Reset, "\n\n    ", ctc.ForegroundBrightCyan, "CPLY_TOKEN", ctc.Reset, ctc.ForegroundBrightCyan, "CPLY_FIRST", ctc.Reset, ctc.ForegroundBrightCyan, "CPLY_PRINT", ctc.Reset, ctc.ForegroundBrightCyan, "CPLY_NOCOLOR", ctc.Reset, "\n\nCPLY v0.2 by ASVVVAD (", ctc.ForegroundBrightCyan, "https://asvvvad.eu.org", ctc.Reset, ") | Github: ", ctc.ForegroundBrightCyan, "(https://github.com/asvvvad/cply)", ctc.Reset)
	// parse flags
	var chooseFirst bool
	flag.BoolVar(&chooseFirst, "1", false, "")
	flag.BoolVar(&chooseFirst, "first", false, "")
	var printLyrics bool
	flag.BoolVar(&printLyrics, "p", false, "")
	flag.BoolVar(&printLyrics, "print", false, "")
	var noColors bool
	flag.BoolVar(&noColors, "n", false, "")
	flag.BoolVar(&noColors, "no-color", false, "")
	var showHelp bool
	flag.BoolVar(&showHelp, "h", false, "")
	flag.BoolVar(&showHelp, "help", false, "")

	flag.Parse()

	if showHelp {
		fmt.Println(help)
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		err("Song not specified specified\n", false)
		fmt.Println(help)
		os.Exit(1)
	}

	// the search query
	var q string

	q = strings.TrimSpace(strings.Join(flag.Args(), " "))

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


	token := os.Getenv("CPLY_TOKEN")
	client := genius.NewClient(token)
	res, _ := client.Search.Get(q)

	if len(res) == 0 {
		err("No results found, try again.", true)
		os.Exit(1)
	}

	var url string
	var title string

	if chooseFirst {
		url = res[0].Result.URL
		title = res[0].Result.FullTitle
	} else {
		info("Found results for \"" + q + "\":")
		for i := 0; i < len(res); i++ {
			fmt.Println(ctc.ForegroundBrightCyan, i, ctc.Reset, res[i].Result.FullTitle)
		}

		selected := -1

		for selected < 0 || selected >= len(res) {
			fmt.Print("\n > Choose a song [0]: ")
			fmt.Scanln(&selected)
			if (selected == -1) {
				selected = 0
			}
		}


		url = res[selected].Result.URL
		title = res[selected].Result.FullTitle

		info("Lyrics for "+title+" are being fetched...")
	}

	l, _ := lyrics.LyricsURL(url)
	lyrics := strings.ReplaceAll(fmt.Sprint(title, "\n", strings.Join(l, "\n")), "[", "\n[")

	if clipboard.Unsupported {
		err("Can't copy to the clipboard. Here are the lyrics:", false)
		if !noColors {
			lyrics = strings.ReplaceAll(lyrics, "[", fmt.Sprint(ctc.ForegroundBrightYellow, "["))
			lyrics = strings.ReplaceAll(lyrics, "]", fmt.Sprint("]", ctc.Reset))
		}
		fmt.Println(lyrics)
		os.Exit(1)
	} else {
		clipboard.WriteAll(lyrics)
		if !(printLyrics && chooseFirst) {
			success("Lyrics copied successfully!")
			if printLyrics {
				success("Printing lyrics")
			}
		}
		if printLyrics {
			if !noColors {
				lyrics = strings.ReplaceAll(lyrics, "[", fmt.Sprint(ctc.ForegroundBrightYellow, "["))
				lyrics = strings.ReplaceAll(lyrics, "]", fmt.Sprint("]", ctc.Reset))
			}
			fmt.Println(lyrics)
		} 

		os.Exit(0)
	}

}
