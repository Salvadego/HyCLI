package config

import (
	"HyCLI/internal/config"
	"HyCLI/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	ConfigSelectCmd.Flags().StringVarP(&client, "client", "c", "", "client name")

	ConfigSelectCmd.MarkFlagRequired("client")
	ConfigSelectCmd.RegisterFlagCompletionFunc("client", utils.ClientNameCompletion)
}

var ConfigSelectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select active client",
	Run: func(cmd *cobra.Command, args []string) {
		name := client
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Printf("failed to load config: %s\n", err.Error())
			return
		}
		if _, ok := cfg.Clients[name]; !ok {
			fmt.Printf("client %q not found", name)
			return
		}

		cfg.DefaultClient = name
		err = config.SaveConfig(cfg)
		if err != nil {
			fmt.Printf("failed to save config: %s\n", err.Error())
			return
		}
	},
}
