package utils

import (
	"strings"

	"github.com/spf13/cobra"
)

func CompleteScriptType(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var completions []string
	for _, f := range []string{"groovy", "beanshell", "javascript"} {
		if strings.HasPrefix(f, toComplete) {
			completions = append(completions, f)
		}
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

func CompleteOutputFormat(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var completions []string
	for _, f := range []string{"table", "json"} {
		if strings.HasPrefix(f, toComplete) {
			completions = append(completions, f)
		}
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}
