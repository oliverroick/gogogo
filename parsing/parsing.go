package main

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {
	var feeds [2]string
	feeds[0] = "https://macwright.org/rss.xml"
	feeds[1] = "http://www.zeldman.com/feed/"

	var aWeekAgo = time.Now().Add(-21 * 24 * time.Hour)
	fmt.Println(aWeekAgo)

	fp := gofeed.NewParser()

	for _, url := range feeds {
		feed, _ := fp.ParseURL(url)
		fmt.Println(feed.Title)
		for _, item := range feed.Items {
			var published = time.Now()
			if item.PublishedParsed != nil {
				published = *item.PublishedParsed
			}
			if aWeekAgo.Before(published) {
				fmt.Println(item.Title)
			}
		}
	}
}
