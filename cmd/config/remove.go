package config

import (
	"HyCLI/internal/config"
	"HyCLI/internal/utils"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	ConfigRemoveCmd.Flags().StringVarP(&client, "client", "c", "", "client name")
	ConfigRemoveCmd.RegisterFlagCompletionFunc("client", utils.ClientNameCompletion)
}

var ConfigRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a client configuration",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Header("Removing client")

		cfg, err := config.LoadConfig()
		if err != nil {
			utils.Error("failed to load config: %s\n", err.Error())
			return
		}

		if len(cfg.Clients) == 0 {
			utils.Info("no clients configured")
			return
		}

		name := client
		if name == "" {
			utils.Info("Available clients:")
			for n := range cfg.Clients {
				fmt.Printf("  %s\n", n)
			}
			fmt.Print("Enter client name to remove: ")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			name = strings.TrimSpace(input)
		}

		if _, ok := cfg.Clients[name]; !ok {
			utils.Error("client %q not found", name)
			return
		}

		fmt.Printf("%sAre you sure you want to delete '%s'? [y/N]%s ", utils.ColorBold, name, utils.ColorReset)
		reader := bufio.NewReader(os.Stdin)
		confirm, _ := reader.ReadString('\n')
		if !strings.HasPrefix(strings.ToLower(confirm), "y") {
			utils.Info("Operation cancelled.")
			return
		}

		delete(cfg.Clients, name)
		if cfg.DefaultClient == name {
			cfg.DefaultClient = ""
		}

		err = config.SaveConfig(cfg)
		if err != nil {
			utils.Error("failed to save config: %s\n", err.Error())
			return
		}
		utils.Success("Client %s removed successfully\n", name)
	},
}
