package client

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Vivena/Toy-QKD-IKE/gateway/networking"
)

// Program name
const ProgramName = "client"

func Cmd() *cobra.Command {
	return cobraCommand
}

var cobraCommand = &cobra.Command{
	Use:   "start",
	Short: "Start the vpn client.",
	Long:  `Start the vpn client.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}
		// Parsing of the command line is done so silence cmd usage
		cmd.SilenceUsage = true

		// TODO: error handeling
		qkd_addr, _ := cmd.Flags().GetString("qkd-ip")
		qkd_port, _ := cmd.Flags().GetString("qkd-port")
		sa_ip, _ := cmd.Flags().GetString("sa-ip")

		fmt.Print("Starting the server")
		client := networking.NewClient(sa_ip, qkd_port, qkd_addr)

		client.Init_IKE_SA()
		
		return nil
	},
}
