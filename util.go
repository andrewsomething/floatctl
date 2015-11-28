package main

import (
	"fmt"
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
	client := godo.NewClient(oauthClient)
	client.BaseURL = GodoBase
	return client
}

// WhoAmI returns a Droplet's ID via the Metadata service if run on one.
func WhoAmI(cmd *cobra.Command) int {
	client := metadata.NewClient(metadata.WithBaseURL(MetadataBase))

	id, err := client.DropletID()
	if err != nil {
		fmt.Println("Error: ", err)
		cmd.Help()
		os.Exit(1)
	}

	return id
}

func AssignedFIP(cmd *cobra.Command) string {
	meta := metadata.NewClient(metadata.WithBaseURL(MetadataBase))
	all, err := meta.Metadata()

	if err != nil {
		fmt.Println("Error: ", err)
		cmd.Help()
		os.Exit(1)
	}

	fip := all.FloatingIP.IPv4.IPAddress

	return fip
}
