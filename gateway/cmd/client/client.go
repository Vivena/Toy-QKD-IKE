package client

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/Vivena/Toy-QKD-IKE/gateway/networking"
)

// This is the "client" side of the implementation.
// It is used to emulate a VPN end-point running a key generation, and rekey
// This is supposed to be integrated in the VPN itself at a later date, and is here
// only to illustrate the communication involved in IKEv2

// Program name
const ProgramName = "vpn"

func Cmd() *cobra.Command {
	return cobraCommand
}

var cobraCommand = &cobra.Command{
	Use:   "client",
	Short: "Start the vpn client.",
	Long:  `Start the vpn client.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}
		// Parsing of the command line is done so silence cmd usage
		// TODO: handel multiple level of verbose, and do proper logging
		cmd.SilenceUsage = true

		// TODO: error handeling
		qkd_addr, _ := cmd.Flags().GetString("qkd-ip")
		qkd_port, _ := cmd.Flags().GetString("qkd-port")
		sa_ip, _ := cmd.Flags().GetString("sa-ip")

		fmt.Println("Starting the client")
		// create a new client with the informations provided
		client, err := networking.NewCli(sa_ip, qkd_addr, qkd_port)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Initializing the SA")
		// run the Init_IKE_SA
		client.Init_IKE_SA()
		// TODO: run the IKE_AUTH
		// TODO: run the rekey

		return nil
	},
}
