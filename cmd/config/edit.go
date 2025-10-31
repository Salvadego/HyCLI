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
	ConfigEditCmd.Flags().StringVarP(&address, "address", "a", "", "address of the server")
	ConfigEditCmd.Flags().StringVarP(&user, "user", "u", "", "username")
	ConfigEditCmd.Flags().StringVarP(&pass, "password", "p", "", "password")
	ConfigEditCmd.Flags().StringVarP(&client, "client", "c", "", "client name")
	ConfigEditCmd.RegisterFlagCompletionFunc("client", utils.ClientNameCompletion)
}

var ConfigEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing client configuration",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Header("Edit Client Configuration")

		cfg, err := config.LoadConfig()
		if err != nil {
			utils.Error("failed to load config: %s", err.Error())
			return
		}

		name := client
		reader := bufio.NewReader(os.Stdin)
		if name == "" {
			utils.Info("Available clients:")
			for n := range cfg.Clients {
				fmt.Printf("  %s\n", n)
			}
			fmt.Print("Enter client to edit: ")
			input, _ := reader.ReadString('\n')
			name = strings.TrimSpace(input)
		}

		c, ok := cfg.Clients[name]
		if !ok {
			utils.Error("client %q not found", name)
			return
		}

		fmt.Printf("%sEditing client:%s %s\n", utils.ColorBold, utils.ColorReset, name)

		if address == "" {
			fmt.Printf("Address [%s]: ", c.Address)
			input, _ := reader.ReadString('\n')
			address = strings.TrimSpace(input)
			if address == "" {
				address = c.Address
			}
		}
		if user == "" {
			fmt.Printf("Username [%s]: ", c.User)
			input, _ := reader.ReadString('\n')
			user = strings.TrimSpace(input)
			if user == "" {
				user = c.User
			}
		}
		if pass == "" {
			fmt.Print("Password [hidden, leave empty to keep current]: ")
			input, _ := reader.ReadString('\n')
			pass = strings.TrimSpace(input)
			if pass == "" {
				pass = c.Password
			}
		}

		cfg.Clients[name] = config.ClientEntry{Address: address, User: user, Password: pass}
		err = config.SaveConfig(cfg)
		if err != nil {
			utils.Error("failed to save config: %s", err.Error())
			return
		}

		utils.Success("Client %s updated successfully", name)
	},
}
