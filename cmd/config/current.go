package config

import (
	"HyCLI/internal/config"
	"HyCLI/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var ConfigCurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show current active client configuration",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Header("Current Active Configuration")

		cfg, err := config.LoadConfig()
		if err != nil {
			utils.Error("failed to load config: %s", err.Error())
			return
		}

		if cfg.DefaultClient == "" {
			utils.Info("No active client configuration.")
			return
		}

		c, ok := cfg.Clients[cfg.DefaultClient]
		if !ok {
			utils.Error("Active client not found in config.")
			return
		}

		fmt.Printf("%sActive client:%s %s\n", utils.ColorBold, utils.ColorReset, cfg.DefaultClient)
		fmt.Printf("  URL: %s\n", c.Address)
		fmt.Printf("  User: %s\n", c.User)
		utils.Success("Configuration loaded successfully.")
	},
}
