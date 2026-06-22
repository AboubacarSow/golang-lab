package main

import (
	"time"

	"github.com/AboubacarSow/golang-lab/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type user struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uuid.UUID `json:"user_id"`
}
type follows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedId    uuid.UUID `json:"feed_id"`
	UserId    uuid.UUID `json:"user_id"`
}

func toUserDto(databaseUser database.User) user {
	return user{
		ID:        databaseUser.ID,
		CreatedAt: databaseUser.CreatedAt,
		UpdatedAt: databaseUser.UpdatedAt,
		Name:      databaseUser.Name.String,
		ApiKey:    databaseUser.ApiKey,
	}
}

func toFeedDto(databaseFeed database.Feed) feed {
	return feed{
		ID:        databaseFeed.ID,
		CreatedAt: databaseFeed.CreatedAt,
		UpdatedAt: databaseFeed.UpdatedAt,
		Name:      databaseFeed.Name.String,
		Url:       databaseFeed.Url,
		UserId:    databaseFeed.UserID,
	}
}

func toFeedDtos(Feeds []database.Feed) []feed {
	feeds := []feed{}

	for _, feed := range Feeds {
		feedDto := toFeedDto(feed)
		feeds = append(feeds, feedDto)
	}
	return feeds
}

func toFollowDto(databaseFeedFollow database.FeedFollow) follows {
	return follows{
		ID:        databaseFeedFollow.ID,
		CreatedAt: databaseFeedFollow.CreatedAt,
		UpdatedAt: databaseFeedFollow.UpdatedAt,
		UserId:    databaseFeedFollow.UserID,
		FeedId:    databaseFeedFollow.FeedID,
	}
}

func toFollowsDtos(FeedFollows []database.FeedFollow) []follows {
	follows := []follows{}

	for _, follow := range FeedFollows {
		feedFollowDto := toFollowDto(follow)
		follows = append(follows, feedFollowDto)
	}
	return follows
}

//func toDtos(databaseUsers []database.User) []user {
//	users := []user{}

//	for _, databaseUser := range databaseUsers {
//		user := toDto(databaseUser)
//		users = append(users, user)
//	}

//	return users
//}
