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

var (
	address string
	user    string
	pass    string
	client  string
)

func init() {
	ConfigAddCmd.Flags().StringVarP(&address, "address", "a", "", "address of the server")
	ConfigAddCmd.Flags().StringVarP(&user, "user", "u", "", "username")
	ConfigAddCmd.Flags().StringVarP(&pass, "password", "p", "", "password")
	ConfigAddCmd.Flags().StringVarP(&client, "client", "c", "", "client name")
}

var ConfigAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new client configuration",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Header("Add New Client Configuration")

		reader := bufio.NewReader(os.Stdin)
		if client == "" {
			fmt.Print("Client name: ")
			input, _ := reader.ReadString('\n')
			client = strings.TrimSpace(input)
		}
		if address == "" {
			fmt.Print("HAC URL: ")
			input, _ := reader.ReadString('\n')
			address = strings.TrimSpace(input)
		}
		if user == "" {
			fmt.Print("Username: ")
			input, _ := reader.ReadString('\n')
			user = strings.TrimSpace(input)
		}
		if pass == "" {
			fmt.Print("Password: ")
			input, _ := reader.ReadString('\n')
			pass = strings.TrimSpace(input)
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			utils.Error("failed to load config: %s", err.Error())
			return
		}

		if _, exists := cfg.Clients[client]; exists {
			utils.Error("Client %q already exists.", client)
			return
		}

		cfg.Clients[client] = config.ClientEntry{Address: address, User: user, Password: pass}
		err = config.SaveConfig(cfg)
		if err != nil {
			utils.Error("failed to save config: %s", err.Error())
			return
		}

		utils.Success("Client %s added successfully", client)
	},
}
