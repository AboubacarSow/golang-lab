package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/AboubacarSow/golang-lab/rss-aggregator/internal/database"
)

func scraperSpinner(db *database.Queries, concurrent int, interval time.Duration) {
	//fetch feeds
	log.Printf("Scraping on %v goroutines every %v duration\n", concurrent, interval)
	timer := time.NewTicker(interval)
	for ; ; <-timer.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrent))
		if err != nil {
			log.Println("Failed to fetch feeds from database" + err.Error())
			continue
		}
		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go handler(db, feed, wg)
		}
		wg.Wait()
	}

}

func handler(database *database.Queries, feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := database.MarkAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Failed to make as fetch feed with Id:%v\n", feed.ID)
		return
	}

	rssFeed, err := getFeedFromUrl(feed.Url)
	if err != nil {
		log.Printf("Error occured while fetching rss feed:%v\n", err.Error())
		return
	}
	for _, item := range rssFeed.Channel.Item {
		log.Printf("- Found post:%v\n", item.Title)
	}

	log.Printf("Total post for feed with Id-%v:%v\n", feed.ID, len(rssFeed.Channel.Item))

}
