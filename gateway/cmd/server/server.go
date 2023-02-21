package server

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Vivena/Toy-QKD-IKE/gateway/networking"
)

// Program name
const ProgramName = "vpn"

func Cmd() *cobra.Command {
	return cobraCommand
}

var cobraCommand = &cobra.Command{
	Use:   "server",
	Short: "Start the vpn server.",
	Long:  `Start the vpn server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}
		// Parsing of the command line is done so silence cmd usage
		cmd.SilenceUsage = true

		//TODO: add the flags so we can set the infos for the qkd
		// qkd_addr, _ := cmd.Flags().GetString("qkd-ip")
		// qkd_port, _ := cmd.Flags().GetString("qkd-port")

		fmt.Println("Starting the server")
		// create a new server with default configuration
		server, err := networking.NewServ()
		if err != nil {
			return err
		}

		// TODO: once we incorporate the client and server in the same instance,
		// we will run the server in a go routine.
		// Furthemore, we will also run another server to receive uapi calls so we can controle the vpn,
		// as well as another go routine running the TUN

		// start the server
		server.Start()
		return nil
	},
}
