package utils

import (
	"strings"

	"github.com/spf13/cobra"
)

func CompleteOutputFormat(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var completions []string
	for _, o := range []string{"table", "json", "csv"} {
		if strings.HasPrefix(o, toComplete) {
			completions = append(completions, o)
		}
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}
