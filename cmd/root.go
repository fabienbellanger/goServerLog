package cmd

import (
	"github.com/fabienbellanger/goServerLog/shared/toolbox"
	"github.com/spf13/cobra"
)

// RootCommand define the root command
var RootCommand = &cobra.Command{
	Use:     "goServerLog",
	Short:   "goServerLog manage logs from Nginx and applications",
	Long:    "goServerLog manage logs from Nginx and applications",
	Version: "0.0.1",
}

var port, defaultPort int

// Execute starts Cobra
func Execute() {
	// Lancement de la commande racine
	// -------------------------------
	err := RootCommand.Execute()
	toolbox.CheckError(err, 1)
}
