package utils

import (
	"os"
	"sort"
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

func CompleteTemplateName(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	scriptDir := os.Getenv("HYCLI_SCRIPT_HOME")
	if scriptDir == "" {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	files, err := os.ReadDir(scriptDir)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	var completions []string
	for _, f := range files {
		if strings.HasPrefix(f.Name(), toComplete) {
			completions = append(completions, f.Name())
		}
	}
	sort.Strings(completions)
	return completions, cobra.ShellCompDirectiveNoFileComp
}
