package main

import (
	"fmt"
	"os"

	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

// Show information about a Floating IP.
func Show(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		floatingIP := doShow(args[0])
		printShow(floatingIP)
	} else if len(args) == 0 {
		fip := AssignedFIP(cmd)
		floatingIP := doShow(fip)
		printShow(floatingIP)
	} else {
		cmd.Help()
	}
}

func doShow(fip string) *godo.FloatingIP {
	client := GetClient(Token)
	floatingIP, _, err := client.FloatingIPs.Get(fip)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return floatingIP
}

func printShow(fip *godo.FloatingIP) {
	fmt.Println("Floating IP\tRegion\t\tDroplet ID\tDroplet Name")
	fmt.Println("-----------\t------\t\t----------\t------------")

	ip := fip.IP
	region := fip.Region.Name
	if fip.Droplet != nil {
		dropletID := fip.Droplet.ID
		dropletName := fip.Droplet.Name
		fmt.Printf("%v\t%v\t%v\t\t%v\n", ip, region, dropletID, dropletName)
	} else {
		fmt.Printf("%v\t%v\n", ip, region)
	}
}
