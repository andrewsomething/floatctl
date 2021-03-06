package main

import (
	"fmt"
	"os"

	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

// Unassign a Floating IP from a Droplet.
func Unassign(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		doUnassign(args[0])
	} else if len(args) == 0 {
		fip := AssignedFIP(cmd)
		action := doUnassign(fip)

		fmt.Printf("Unassigning %v in %v...\n", fip, action.Region.Slug)
	} else {
		cmd.Help()
	}
}

func doUnassign(fip string) *godo.Action {
	client := GetClient(Token)
	action, _, err := client.FloatingIPActions.Unassign(fip)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return action
}
