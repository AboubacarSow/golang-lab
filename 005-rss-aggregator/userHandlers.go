package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/AboubacarSow/golang-lab/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type parameters struct {
	Name string `json:"name"`
}

func (apiconf ApiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {

	decorator := json.NewDecoder(r.Body)
	params := parameters{}
	err := decorator.Decode(&params)
	if err != nil {
		errorHelper(w, 400, "Error occured while decoding data")
		return
	}

	new_user, err := apiconf.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Name: sql.NullString{
			String: params.Name,
			Valid:  true,
		},
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		errorHelper(w, 500, fmt.Sprintf("Error occured while creating user with name:%v", params.Name))
		return
	}
	jsonHelper(w, 201, toUserDto(new_user))
}

func getUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	jsonHelper(w, 200, toUserDto(user))
}

//func (apiconf ApiConfig) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {

//	users, err := apiconf.DB.GetAll(r.Context())

//	if err != nil {
//		errorHelper(w, 400, fmt.Sprintf("Error while fetching users:%v", err))
//		return
//	}
//	jsonHelper(w, 200, toDtos(users))
//}

func (apiconf ApiConfig) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	requestParams := strings.TrimPrefix(r.URL.Path, "/api/v1/users/")

	if requestParams == "" {
		errorHelper(w, 400, "Error:id param cannot be empty or nil")
		return
	}
	id := uuid.MustParse(requestParams)
	err := apiconf.DB.DeleteUser(r.Context(), id)

	if err != nil {
		errorHelper(w, 400, fmt.Sprintf("Error while getting user with id=%s:%v", id, err))
		return
	}
	jsonHelper(w, 204, struct{}{})
}
