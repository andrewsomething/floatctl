package main

import (
	"os"

	"github.com/digitalocean/go-metadata"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
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

// If on a Droplet, returns its ID via the Metadata service.
func WhoAmI(cmd *cobra.Command) int {
	client := metadata.NewClient()

	id, err := client.DropletID()
	if err != nil {
		cmd.Help()
		os.Exit(1)
	}

	return id
}
