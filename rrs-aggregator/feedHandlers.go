package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AboubacarSow/golang-lab/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type createFeedDto struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (apiconf ApiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request,user database.User) {
	
	decoder := json.NewDecoder(r.Body)
	requestDto := createFeedDto{}
	err := decoder.Decode(&requestDto)

	if err != nil {
		errorHelper(w, 400, "Error occured while decoding data")
		return
	}
	feed, err := apiconf.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		Name: sql.NullString{
			String: requestDto.Name,
			Valid:  true,
		},
		Url:       requestDto.Url,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
	})

	if err != nil {
		errorHelper(w, 500, "Failed to create feed."+err.Error())
		return
	}

	jsonHelper(w, 201, toFeedDto(feed))

}

func (apiconf ApiConfig) getAllFeedsHandler(w http.ResponseWriter, r *http.Request){
	feeds, err := apiconf.DB.GetAllFeeds(r.Context())

	if err != nil {
		errorHelper(w, 400, fmt.Sprintf("Error while fetching feeds:%v", err))
		return
	}
	jsonHelper(w, 200, toFeedDtos(feeds))
}
