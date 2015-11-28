package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

// Assign an existing Floating IP to a Droplet.
func Assign(cmd *cobra.Command, args []string) {
	if len(args) == 2 {
		id, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		act := doAssign(args[0], id)

		fmt.Printf("Assigning %v to Droplet %v in %v...\n", args[0], id, act.Region.Slug)
	} else if len(args) == 1 {
		id := WhoAmI(cmd)
		act := doAssign(args[0], id)

		fmt.Printf("Assigning %v to Droplet %v in %v...\n", args[0], id, act.Region.Slug)

	} else {
		cmd.Help()
	}
}

// Make the actual API call to assign the Floating IP
func doAssign(ip string, id int) *godo.Action {
	client := GetClient(Token)

	action, _, err := client.FloatingIPActions.Assign(ip, id)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return action
}
