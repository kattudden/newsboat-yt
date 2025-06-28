package main

import (
	"flag"
	"fmt"
	"kattudden/newsboat-yt/database"
	"kattudden/newsboat-yt/download"
)

func main() {

	newUrl := flag.String("u", "", "add url to download queue.")
	processQueue := flag.Bool("d", false, "start downloading from queue.")

	flag.Parse()

	if *newUrl != "" {
		fmt.Println("Adding new URL.")
		database.InsertUrl(*newUrl)
		return
	}

	if *processQueue {
		fmt.Println("Processing queue...")

		urls := database.GetUrls()
		for _, url := range urls {
			err := download.YoutubeVideo(url.Link)
			if err != nil {
				fmt.Println("failed to download video: ", url.Link)
				fmt.Println("but i don't care and mark it downloaded...")
				database.MarkUrlDownloaded(url.ID)
			}

			fmt.Println("successfully downloaded: ", url.Link)
			database.MarkUrlDownloaded(url.ID)
		}

		fmt.Println("Done!")
		return
	}

	fmt.Println("Missing Argmuent; doing nothing.")
}
