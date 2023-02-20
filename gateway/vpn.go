package main

import (
	"os"
	"strings"

	"github.com/Vivena/Toy-QKD-IKE/gateway/cmd/client"
	"github.com/Vivena/Toy-QKD-IKE/gateway/cmd/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var mainCmd = &cobra.Command{Use: "vpn"}

func main() {
	// For environment variables.
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	mainCmd.AddCommand(client.Cmd())
	mainCmd.AddCommand(server.Cmd())

	client.Cmd().PersistentFlags().String("qkd-ip", "", "IP for the QKD REST API.")
	client.Cmd().PersistentFlags().String("qkd-port", "", "Port for the QKD REST API.")
	client.Cmd().PersistentFlags().String("sa-ip", "", "IP for the SA we want to connect to.")

	if mainCmd.Execute() != nil {
		os.Exit(1)
	}
}
