package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"hycli-script/internal"
	"hycli-script/utils"

	"github.com/Salvadego/hac/hac"
	"github.com/spf13/cobra"
)

var (
	scriptFile string
	scriptOut  string
	scriptType string
)

var scriptCmd = &cobra.Command{
	Use:   "run [SCRIPT]",
	Short: "Execute a script",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var scriptContent string
		if scriptFile != "" {
			b, err := os.ReadFile(scriptFile)
			if err != nil {
				return err
			}
			scriptContent = string(b)
		} else {
			scriptContent = strings.Join(args, "\n")
			if scriptContent == "" {
				return fmt.Errorf("no script provided")
			}
		}

		c := internal.New(baseURL, username, password, skipVerify)
		ctx := context.Background()
		if err := internal.Login(c, ctx); err != nil {
			return err
		}

		if scriptType == "" {
			scriptType = "groovy"
		}
		var stype hac.ScriptType
		stype = hac.ScriptType(scriptType)

		resp, err := c.Groovy.Execute(ctx, hac.GroovyRequest{
			Script:     scriptContent,
			ScriptType: stype,
		})
		if err != nil {
			return err
		}

		pritnOutput(resp, internal.OutputFormat(scriptOut))

		return nil
	},
}

func pritnOutput(resp *hac.GroovyResponse, format internal.OutputFormat) {
	switch format {
	case internal.OutputTable:
		if resp.Output != "" {
			fmt.Println(resp.Output)
		}
		if resp.Result != "" {
			fmt.Println(resp.Result)
		}
		if resp.Stacktrace != "" {
			fmt.Println(resp.Stacktrace)
		}
	case internal.OutputJSON:
		b, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	}
}

func init() {
	rootCmd.AddCommand(scriptCmd)
	scriptCmd.Flags().StringVarP(&scriptFile, "file", "f", "", "File containing the script")
	scriptCmd.Flags().StringVarP(&scriptType, "type", "t", "", "Script type: groovy|javascript|beanshell")
	scriptCmd.Flags().StringVarP(&scriptOut, "output", "o", "table", "table|json")

	scriptCmd.RegisterFlagCompletionFunc("output", utils.CompleteOutputFormat)
	scriptCmd.RegisterFlagCompletionFunc("type", utils.CompleteScriptType)
}
