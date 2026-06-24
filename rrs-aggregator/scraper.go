package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/AboubacarSow/golang-lab/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func scraperSpinner(db *database.Queries, concurrent int, interval time.Duration) {
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

func handler(db *database.Queries, feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := db.MarkAsFetched(context.Background(), feed.ID)
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
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubat, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			log.Printf("Failed to parse published date for item with title:%s.Error Message:%v", item.Title, err.Error())
			continue
		}
		err = db.AddPost(context.Background(), database.AddPostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubat,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				continue
			}
			log.Printf("Failed to create post for feed:%v..Erro message:%v", feed.Name.String, err.Error())
			continue
		}
	}

	log.Printf("fetch feed with name:%s\n", feed.Name.String)

}
