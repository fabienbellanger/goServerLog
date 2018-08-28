package cmd

import (
	"fmt"

	"github.com/fabienbellanger/goServerLog/client"
	"github.com/fabienbellanger/goServerLog/shared/toolbox"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	// Flag
	// ----
	defaultPort = 3000
	ClientCommand.Flags().IntVarP(&port, "port", "p", defaultPort, "listened port")

	// Ajout de la commande à la commande racine
	RootCommand.AddCommand(ClientCommand)
}

// ClientCommand : Client command
var ClientCommand = &cobra.Command{
	Use:   "client",
	Short: "Start RPC client",

	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|------------------|
|                  |
| Start RPC client |
|                  |
|------------------|

		`)

		// Test du port
		// ------------
		if port < 1000 || port > 10000 {
			port = defaultPort
		}
		fmt.Print("Client listening on port ")
		color.Green(toolbox.IntToString(port) + "\n\n")

		// Lancement du client RPC
		// -----------------------
		// client.Listen(port)
		client.GetLogs()
	},
}
