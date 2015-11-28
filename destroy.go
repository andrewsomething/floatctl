package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
