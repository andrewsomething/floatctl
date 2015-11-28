package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
