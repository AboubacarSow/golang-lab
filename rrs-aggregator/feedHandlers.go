package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/AboubacarSow/golang-lab/rss-aggregator/internal/auth"
	"github.com/AboubacarSow/golang-lab/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type createFeedDto struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (apiconf ApiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request) {
	apikey, err := auth.GetApiKey(r.Header)

	if err != nil {
		errorHelper(w, 401, err.Error())
		return
	}

	user, err := apiconf.DB.GetUserByKey(r.Context(), apikey)
	if err != nil {

		errorHelper(w, 400, err.Error())
	}
	decoder := json.NewDecoder(r.Body)
	requestDto := createFeedDto{}
	err = decoder.Decode(&requestDto)

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
