package pokeapi

import (
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

type Client struct {
	httpClient http.Client
	Cache      *pokecache.Cache
}

const baseUrl = "https://pokeapi.co/api/v2"

func NewClient(interval time.Duration) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		Cache: pokecache.New(interval),
	}
}
