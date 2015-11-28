package main

import (
	"fmt"
	"os"

	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

// List information about existing Floating IPs.
func List(cmd *cobra.Command, args []string) {
	client := GetClient(Token)

	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}

	floatingIPs, _, err := client.FloatingIPs.List(opt)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Floating IP\tRegion\t\tDroplet ID\tDroplet Name")
	fmt.Println("-----------\t------\t\t----------\t------------")

	for i := range floatingIPs {
		ip := floatingIPs[i].IP
		region := floatingIPs[i].Region.Name
		if floatingIPs[i].Droplet != nil {
			dropletID := floatingIPs[i].Droplet.ID
			dropletName := floatingIPs[i].Droplet.Name
			fmt.Printf("%v\t%v\t%v\t\t%v\n", ip, region, dropletID, dropletName)
		} else {
			fmt.Printf("%v\t%v\n", ip, region)
		}
	}
}
