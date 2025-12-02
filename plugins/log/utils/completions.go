package utils

import (
	"strings"

	"github.com/spf13/cobra"
)

func CompleteLogLevel(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var completions []string
	for _, level := range []string{
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
		"FATAL",
		"OFF",
		"ALL",
	} {
		if strings.HasPrefix(level, toComplete) {
			completions = append(completions, level)
		}
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}
