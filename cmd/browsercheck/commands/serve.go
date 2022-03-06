package commands

import (
	"fmt"

	"github.com/kironono/browsercheck/internal/config"
	"github.com/kironono/browsercheck/internal/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.RunE = serverCommandF
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the server",
	RunE:  serverCommandF,
}

func serverCommandF(command *cobra.Command, args []string) error {
	fc := config.NewFileConfig(getConfigPath(command))
	conf, err := fc.Load()

	if err != nil {
		return fmt.Errorf("failed to load configuration\n%v\n", err)
	}

	return server.Run(conf)
}
