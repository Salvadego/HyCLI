package cmd

import (
	"HyCLI/cmd/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage hycli client configurations",
	Long:  "Create, list, select, and modify hycli client configurations in YAML format",
}

func init() {
	configCmd.AddCommand(config.ConfigListCmd)
	configCmd.AddCommand(config.ConfigSelectCmd)
	configCmd.AddCommand(config.ConfigAddCmd)
	configCmd.AddCommand(config.ConfigEditCmd)
	configCmd.AddCommand(config.ConfigRemoveCmd)
	configCmd.AddCommand(config.ConfigCurrentCmd)
	rootCmd.AddCommand(configCmd)
}
