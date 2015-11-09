package main

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	Token   string
	Droplet string
	Region  string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "floatctl",
		Short: "Control DigitalOcean Floatin IPs.",
	}

	var cmdShow = &cobra.Command{
		Use:   "show [Floating IP]",
		Short: "Show information about a Floating IP",
		Long: `Show information about a Floating IP.

If run on a DigitalOcean Droplet and a Floating IP is assigned to
it, the argument can be obmitted and it will be inferred using the
Metadata service.`,
		PersistentPreRun: TokenCheck,
		Run:              Show,
	}

	var cmdCreate = &cobra.Command{
		Use:   "create [Droplet ID] --region [Region]",
		Short: "Create a Floating IP",
		Long: `Creates a new Floating IP either assigned to a Droplet or reserved
to a region. If assigning to a specifc Droplet, the '--region' flag
is not needed and will be ignored.

The new Floating IP is returned as the only output making this
command suitable for scripting.`,
		PersistentPreRun: TokenCheck,
		Run:              Create,
	}

	var cmdAssign = &cobra.Command{
		Use:   "assign [Floating IP] [Droplet ID]",
		Short: "Assign a Floating IP to a Droplet",
		Long: `Assigns a Floating IP to a Droplet.

If run on a DigitalOcean Droplet, the Droplet ID argument can be
omitted. If so, the Floating IP will be assigned to the Droplet
itself as reported by the Metadata service.`,
		PersistentPreRun: TokenCheck,
		Run:              Assign,
	}

	var cmdUnassign = &cobra.Command{
		Use:   "unassign [Floating IP]",
		Short: "Unassign a Floating IP",
		Long: `Unassign a Floating IP.

If run on a DigitalOcean Droplet with Floating IP, the argument
can be omitted and the Floating IP will be inferred using the
Metadata service and unassigned.`,
		PersistentPreRun: TokenCheck,
		Run:              Unassign,
	}

	var cmdList = &cobra.Command{
		Use:              "list",
		Short:            "List available Floating IPs",
		PersistentPreRun: TokenCheck,
		Run:              List,
	}

	var cmdDestroy = &cobra.Command{
		Use:              "destroy [Floating IP]",
		Short:            "Destroy a Floatin IP",
		PersistentPreRun: TokenCheck,
		Run:              Destroy,
	}

	var cmdAssigned = &cobra.Command{
		Use:   "assigned",
		Short: "Check if a Floating IP is assigned",
		Long: `Check if a Floating IP is assigned.

This command is meant to be run on a DigitalOcean Droplet. If
there is a Floating IP assigned to the Droplet, its address
will be returned. If not, the command will exit with a non-zero
exit code.`,
		Run: Assigned,
	}

	rootCmd.PersistentFlags().StringVarP(&Token,
		"token", "t", os.Getenv("DIGITALOCEAN_TOKEN"),
		"DigitalOcean API Token - $DIGITALOCEAN_TOKEN",
	)

	rootCmd.AddCommand(cmdCreate, cmdShow, cmdAssign, cmdUnassign, cmdList, cmdDestroy, cmdAssigned)
	cmdCreate.Flags().StringVarP(&Region, "region", "r", "", "Region to reserve Floating IP in")
	rootCmd.Execute()
}
