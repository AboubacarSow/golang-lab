package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AboubacarSow/golang-lab/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiconf ApiConfig) createFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	feedInfo := parameters{}
	err := decoder.Decode(&feedInfo)

	if err != nil {
		errorHelper(w, 400, "Error occured while decoding data"+err.Error())
		return
	}
	databaseFeedFollow, err := apiconf.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feedInfo.FeedId,
	})

	if err != nil {
		errorHelper(w, 500, "Failed to create feed-follow."+err.Error())
		return
	}
	jsonHelper(w, 201, toFollowDto(databaseFeedFollow))

}

func (apiconf ApiConfig) getAllFeedFollowsHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiconf.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		errorHelper(w, 400, fmt.Sprintf("Error while fetching Feed Followed:%v", err))
		return
	}
	jsonHelper(w, 200, toFollowsDtos(feedFollows))
}
