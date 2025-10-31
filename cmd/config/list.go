package config

import (
	"HyCLI/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

var ConfigListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available client configurations",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		if len(cfg.Clients) == 0 {
			fmt.Println("No clients configured")
			return nil
		}

		fmt.Println("Configured clients:")
		for name, c := range cfg.Clients {
			active := ""
			if name == cfg.DefaultClient {
				active = " (active)"
			}
			fmt.Printf("  %s%s\n", name, active)
			fmt.Printf("    URL: %s\n", c.Address)
			fmt.Printf("    User: %s\n", c.User)
		}
		return nil
	},
}
