package evepraisal

import "net/http"

// Client is an interface to interact with the evepraisal api
type Client struct {
	client *http.Client
}

// CreateClient creates a new instance of the evepraisal client
func CreateClient(httpClient *http.Client) *Client {
	return &Client{
		client: httpClient,
	}
}
