package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Robert1060/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		Name string `json: "name"`
		URL  string `json: "url"`
	}

	decoder := json.NewDecoder(r.Body)

	parameters := params{}

	err := decoder.Decode(&parameters)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing data %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      parameters.Name,
		Url:       parameters.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't create feed %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))

}

func (apiCfg ApiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could'n fetch feeds %v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}
