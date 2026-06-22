package main

import (
	"net/http"

	"github.com/AboubacarSow/golang-lab/rss-aggregator/internal/auth"
	"github.com/AboubacarSow/golang-lab/rss-aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)


func (apiconf ApiConfig) authMiddleware(handler authHandler) func (http.ResponseWriter, *http.Request){

	return func(w http.ResponseWriter, r *http.Request){
			apikey, err := auth.GetApiKey(r.Header)

		if err != nil {
			errorHelper(w, 401, err.Error())
			return
		}

		user, err := apiconf.DB.GetUserByKey(r.Context(), apikey)
		if err != nil {

			errorHelper(w, 400, err.Error())
		}

		handler(w, r, user)
	}
}
