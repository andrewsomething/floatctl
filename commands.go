package main

import (
	"fmt"
	"os"
	"strconv"

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

// Make the actual API call to assign the Floating IP
func doAssign(ip string, id int) {
	client := GetClient(Token)

	action, _, err := client.FloatingIPActions.Assign(ip, id)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("Assigning %v to Droplet %v in %v...\n", ip, id, action.Region.Slug)
}

// Assign an existing Floating IP to a Droplet.
func Assign(cmd *cobra.Command, args []string) {
	if len(args) == 2 {
		id, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		doAssign(args[0], id)

	} else if len(args) == 1 {
		id := WhoAmI(cmd)
		doAssign(args[0], id)

	} else {
		cmd.Help()
	}
}

// Unassign a Floating IP from a Droplet.
func Unassign(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		doUnassign(args[0])
	} else if len(args) == 0 {
		fip := AssignedFIP(cmd)
		doUnassign(fip)
	} else {
		cmd.Help()
	}
}

func doUnassign(fip string) {
	client := GetClient(Token)
	action, _, err := client.FloatingIPActions.Unassign(fip)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("Unassigning %v in %v...\n", fip, action.Region.Slug)
}

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

// Destroy a Floating IP.
func Destroy(cmd *cobra.Command, args []string) {
	client := GetClient(Token)

	if len(args) == 1 {
		_, err := client.FloatingIPs.Delete(args[0])

		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		fmt.Println("Successfully destroyed Floating IP:", args[0])
	} else {
		cmd.Help()
	}
}

func Assigned(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		cmd.Help()
	} else if len(args) == 0 {
		fip := AssignedFIP(cmd)
		fmt.Println(fip)
	} else {
		os.Exit(1)
	}
}

// TokenCheck checks if a DigitalOcean API Token has been provided.
func TokenCheck(cmd *cobra.Command, args []string) {
	if Token == "" {
		fmt.Println("The '--token flag or $DIGITALOCEAN_TOKEN environmental variable must be set.")
		os.Exit(1)
	}
}
