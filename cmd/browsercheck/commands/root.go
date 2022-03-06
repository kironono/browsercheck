package commands

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "browsercheck",
	Short: "browsercheck",
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "Configuration file to use.")
}

func getConfigPath(command *cobra.Command) string {
	configPath, _ := command.Flags().GetString("config")
	if configPath == "" {
		configPath = "config.yml"
	}
	return configPath
}

func Run(args []string) error {
	rootCmd.SetArgs(args)
	return rootCmd.Execute()
}
