package cmd

import (
	"HyCLI/internal/config"
	"HyCLI/internal/paths"
	"HyCLI/internal/plugins"
	"HyCLI/internal/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hycli",
	Short: "Hybris Administration Console CLI",
	Long:  "hycli is a CLI for interacting with SAP Hybris HAC endpoints",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		dirs, err := paths.Directories()
		if err != nil {
			return fmt.Errorf("failed to get hycli directories %w", err)
		}

		cfg, err := config.InitializeConfig()
		if err != nil {
			return fmt.Errorf("failed to initialize config: %w", err)
		}

		activeClient := cfg.DefaultClient
		if client != "" {
			if _, ok := cfg.Clients[client]; !ok {
				return fmt.Errorf("client %q not found", client)
			}
			activeClient = client
		}

		if activeClient != "" {
			clientCfg := cfg.Clients[activeClient]
			_ = os.Setenv("HYCLI_CONFIG_HOME", dirs.Config)
			_ = os.Setenv("HYCLI_DATA_HOME", dirs.Data)
			_ = os.Setenv("HYCLI_STATE_HOME", dirs.State)
			_ = os.Setenv("HYCLI_CLIENT_NAME", activeClient)
			_ = os.Setenv("HYCLI_CLIENT_URL", clientCfg.Address)
			_ = os.Setenv("HYCLI_CLIENT_USER", clientCfg.User)
			_ = os.Setenv("HYCLI_CLIENT_PASSWORD", clientCfg.Password)
		}

		return nil
	},
}

var client string

func init() {
	rootCmd.PersistentFlags().StringVarP(&client, "client", "c", "", "Name of the client to use")
	rootCmd.RegisterFlagCompletionFunc("client", utils.ClientNameCompletion)

	for _, p := range plugins.Discover() {
		rootCmd.AddCommand(p)
	}

	rootCmd.AddCommand(completionCmd)
}

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate shell completion",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shell := args[0]
		switch shell {
		case "zsh":
			rootCmd.GenZshCompletion(os.Stdout)
		case "bash":
			rootCmd.GenBashCompletion(os.Stdout)
		case "fish":
			rootCmd.GenFishCompletion(os.Stdout, true)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
