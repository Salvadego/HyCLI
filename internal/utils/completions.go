package utils

import (
	"HyCLI/internal/config"
	"strings"

	"github.com/spf13/cobra"
)

func ClientNameCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	var names []string
	for name := range cfg.Clients {
		if strings.HasPrefix(name, toComplete) {
			names = append(names, name)
		}
	}
	return names, cobra.ShellCompDirectiveNoFileComp
}
