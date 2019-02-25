package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmcdole/gofeed"
)

// FeedEvent contains the Lambda event
type FeedEvent struct {
	URL string
}

// FeedItem contains the post title and a link
type FeedItem struct {
	Title string
	Link  string
}

// Feed contains the feed title and its items
type Feed struct {
	Title string
	Items []FeedItem
}

func mapItem(item *gofeed.Item) FeedItem {
	return FeedItem{
		Title: item.Title,
		Link:  item.Link,
	}
}

func mapItems(items []*gofeed.Item, f func(*gofeed.Item) FeedItem) []FeedItem {
	itemsmap := make([]FeedItem, len(items))
	for i, v := range items {
		itemsmap[i] = f(v)
	}
	return itemsmap
}

// HandleEvent parses the feed for the given url and returns feed title and items
func HandleEvent(event FeedEvent) (Feed, error) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(event.URL)

	return Feed{
		Title: feed.Title,
		Items: mapItems(feed.Items, mapItem),
	}, nil
}

func main() {
	lambda.Start(HandleEvent)
}
