package main

import (
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

// Holds an OAuth token.
type TokenSource struct {
	AccessToken string
}

// Returns an OAuth token.
func (t *TokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: t.AccessToken,
	}, nil
}

// Returns a godo client.
func GetClient(token string) *godo.Client {
	tokenSource := &TokenSource{AccessToken: token}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	return godo.NewClient(oauthClient)
}
