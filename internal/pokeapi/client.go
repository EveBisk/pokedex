package pokeapi

import (
	"net/http"
)

// Client -
type Client struct {
	httpClient http.Client
}

// NewClient -
func NewClient() Client {
	return Client{
		httpClient: http.Client{},
	}
}
