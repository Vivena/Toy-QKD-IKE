package client

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/Vivena/Toy-QKD-IKE/gateway/networking"
)

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
		// cmd.SilenceUsage = true
		// TODO: error handeling
		qkd_addr, _ := cmd.Flags().GetString("qkd-ip")
		qkd_port, _ := cmd.Flags().GetString("qkd-port")
		sa_ip, _ := cmd.Flags().GetString("sa-ip")

		fmt.Println("Starting the client")
		fmt.Println("Getting the QKD key")
		client, err := networking.NewCli(sa_ip, qkd_addr, qkd_port)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Initializing the SA")
		client.Init_IKE_SA()

		return nil
	},
}
