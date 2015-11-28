package main

import (
	"fmt"
	"os"

	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

// List information about existing Floating IPs.
func List(cmd *cobra.Command, args []string) {
	floatingIPs := doList()
	printList(floatingIPs)
}

func doList() []godo.FloatingIP {
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

	return floatingIPs
}

func printList(fips []godo.FloatingIP) {
	fmt.Println("Floating IP\tRegion\t\tDroplet ID\tDroplet Name")
	fmt.Println("-----------\t------\t\t----------\t------------")

	for i := range fips {
		ip := fips[i].IP
		region := fips[i].Region.Name
		if fips[i].Droplet != nil {
			dropletID := fips[i].Droplet.ID
			dropletName := fips[i].Droplet.Name
			fmt.Printf("%v\t%v\t%v\t\t%v\n", ip, region, dropletID, dropletName)
		} else {
			fmt.Printf("%v\t%v\n", ip, region)
		}
	}
}
