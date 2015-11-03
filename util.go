package main

import (
	"os"

	"github.com/digitalocean/go-metadata"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// TokenSource holds an OAuth token.
type TokenSource struct {
	AccessToken string
}

// Token returns an OAuth token.
func (t *TokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: t.AccessToken,
	}, nil
}

// GetClient returns a godo client.
func GetClient(token string) *godo.Client {
	tokenSource := &TokenSource{AccessToken: token}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	return godo.NewClient(oauthClient)
}

// WhoAmI returns a Droplet's ID via the Metadata service if run on one.
func WhoAmI(cmd *cobra.Command) int {
	client := metadata.NewClient()

	id, err := client.DropletID()
	if err != nil {
		cmd.Help()
		os.Exit(1)
	}

	return id
}
