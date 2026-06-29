package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func getFeedFromUrl(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)

	if err != nil {
		fmt.Println("Something went wrong while fetching rss content" + err.Error())
		return RSSFeed{}, err
	}
	rssFeed := RSSFeed{}
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("unable to read response body" + err.Error())
		return RSSFeed{}, err
	}
	defer resp.Body.Close()
	xml.Unmarshal(data, &rssFeed)
	return rssFeed, nil
}

//func outputRssFeedHandler(w http.ResponseWriter, r *http.Request){
//	rssFeed, err := getFeedFromUrl("https://www.wagslane.dev/index.xml")
//	if err != nil {
//		errorHelper(w, http.StatusBadRequest,err.Error())
//		return
//	}
//	xmlHelper(w, 200, rssFeed)
//}
