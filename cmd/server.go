package cmd

import (
	"fmt"

	"github.com/fabienbellanger/goServerLog/server"
	"github.com/fabienbellanger/goServerLog/shared/toolbox"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	// Flag
	// ----
	defaultPort = 3000
	ServerCommand.Flags().IntVarP(&port, "port", "p", defaultPort, "listened port")

	// Ajout de la commande à la commande racine
	RootCommand.AddCommand(ServerCommand)
}

// ServerCommand : Server command
var ServerCommand = &cobra.Command{
	Use:   "server",
	Short: "Start RPC server",

	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|------------------|
|                  |
| Start RPC server |
|                  |
|------------------|

		`)

		// Test du port
		// ------------
		if port < 1000 || port > 10000 {
			port = defaultPort
		}
		fmt.Print("Server runing on port ")
		color.Green(toolbox.IntToString(port) + "\n\n")

		server.Start(port)
	},
}
