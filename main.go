package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

func main() {
	var api string
	var link string
	var high bool
	flag.StringVar(&api, "api", "https://pipedapi.kavin.rocks", "specify an API URL and the one defined ignore config.json")
	flag.StringVar(&link, "url", "", "Frontend URL of any Piped or Youtube link.")
	flag.BoolVar(&high, "Q", false, "set it true to download high quality mp4. (experimental, ffmpeg needed)")
	flag.Parse()

	yt := newPiped(api, high)
	if !checkAvail(yt.APIUrl) {
		fmt.Printf("can't connect to API URL: %s\n", yt.APIUrl)
		os.Exit(1)
	}

	uri, err := url.Parse(link)
	if err != nil {
		fmt.Println("provided URL not valid!")
		os.Exit(1)
	}
	params := uri.Query()
	playlist := params.Get("list")
	video := params.Get("v")
	if playlist == "" && video == "" {
		fmt.Println("provided URL not valid!")
		os.Exit(1)
	}
	if video != "" {
		playlist = ""
		yt.DownVideo(video)
	}
	if playlist != "" {
		yt.DownList(playlist)
	}
}
