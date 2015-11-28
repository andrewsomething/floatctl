package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

// Create a new Floating IP.
func Create(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		createRequest := &godo.FloatingIPCreateRequest{
			DropletID: id,
		}

		floatingIP := doCreate(createRequest)
		fmt.Println(floatingIP.IP)
	} else if Region != "" && len(args) < 1 {
		createRequest := &godo.FloatingIPCreateRequest{
			Region: Region,
		}

		floatingIP := doCreate(createRequest)
		fmt.Println(floatingIP.IP)
	} else {
		cmd.Help()
	}
}

func doCreate(req *godo.FloatingIPCreateRequest) *godo.FloatingIP {
	client := GetClient(Token)
	floatingIP, _, err := client.FloatingIPs.Create(req)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return floatingIP
}
