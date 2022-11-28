package helpers

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

// Create an http.Client
// this method will need to be updated if other scopes are needed https://developers.google.com/identity/protocols/oauth2/scopes
func NewHTTPClient(ctx context.Context) (*http.Client, error) {

	c, _, err := transport.NewHTTPClient(ctx,
		option.WithScopes("https://www.googleapis.com/auth/cloud-platform.read-only"))
	if err != nil {
		log.Fatal(err)
	}

	return c, nil

}
