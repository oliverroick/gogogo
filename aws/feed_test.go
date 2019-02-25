package main

import (
	"testing"

	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/assert"
)

func TestMapItem(t *testing.T) {
	item := gofeed.Item{
		Title: "A blog post",
		Link:  "http://example.com/a-blog-post-html",
	}
	var result interface{} = mapItem(&item)
	r, ok := result.(FeedItem)

	assert := assert.New(t)
	assert.True(ok, "Returned result is not of type `FeedItem`")
	assert.Equal(item.Title, r.Title)
	assert.Equal(item.Link, r.Link)
}

func TestMap(t *testing.T) {
	items := []*gofeed.Item{
		{
			Title: "A blog post",
			Link:  "http://example.com/a-blog-post-html",
		},
		{
			Title: "Another blog post",
			Link:  "http://example.com/another-blog-post-html",
		},
	}

	var result interface{} = mapItems(items, mapItem)
	r, ok := result.([]FeedItem)

	assert := assert.New(t)
	assert.True(ok, "Returned result is not a list of `FeedItem`")
	assert.Equal(len(r), len(items), "Result must have same number of elements as input")
}
