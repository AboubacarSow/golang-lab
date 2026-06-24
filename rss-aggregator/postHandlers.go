package main

import (
	"net/http"

	"github.com/AboubacarSow/golang-lab/rss-aggregator/internal/database"
)

func (apiconf ApiConfig) getPostsByUserHandler(w http.ResponseWriter, r *http.Request, user database.User){

	posts, err:= apiconf.DB.GetPostsByUser(r.Context(),database.GetPostsByUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err!=nil{
		errorHelper(w, http.StatusInternalServerError,err.Error())
		return 
	}
	jsonHelper(w,200,toPostDtos(posts))
}