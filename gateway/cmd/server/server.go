package server

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Vivena/Toy-QKD-IKE/gateway/networking"
)

// Program name
const ProgramName = "server"

func Cmd() *cobra.Command {
	return cobraCommand
}

var cobraCommand = &cobra.Command{
	Use:   "start",
	Short: "Start the vpn server.",
	Long:  `Start the vpn server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}
		// Parsing of the command line is done so silence cmd usage
		cmd.SilenceUsage = true
		fmt.Print("Starting the server")
		var server networking.Serv
		server.Start()
		return nil
	},
}
