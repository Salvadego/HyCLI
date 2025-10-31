package cmd

import (
	"context"
	"flex/internal"
	"flex/utils"
	"fmt"
	"os"

	"github.com/Salvadego/hac/hac"
	"github.com/spf13/cobra"
)

type OutputType string

const (
	TableOutput OutputType = "table"
	JSONOutput  OutputType = "json"
	CSVOutput   OutputType = "csv"
)

var (
	flexSQLFile string
	flexSQLMax  int
)

var flexSQLCmd = &cobra.Command{
	Use:   "sql [SQL_QUERY]",
	Short: "Run raw SQL",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var q string
		if flexSQLFile != "" {
			b, err := os.ReadFile(flexSQLFile)
			if err != nil {
				return err
			}
			q = string(b)
		} else {
			if len(args) == 0 {
				fmt.Println("SQL query is required unless --file is used")
				return cmd.Help()
			}
			q = args[0]
		}

		c := internal.New(baseURL, username, password, skipVerify)
		ctx := context.Background()

		if err := internal.Login(c, ctx); err != nil {
			return err
		}

		resp, err := c.Flex.Execute(ctx, hac.FlexQuery{
			SQLQuery: q,
			User:     username,
			MaxCount: flexSQLMax,
		}, nil)
		if err != nil {
			return err
		}

		flexQueryOutput(resp, OutputType(flexOut))
		return nil
	},
}

func init() {
	flexSQLCmd.Flags().StringVarP(&flexSQLFile, "file", "f", "", "File with SQL")
	flexSQLCmd.Flags().IntVarP(&flexSQLMax, "max", "m", 100, "Max results")
	flexSQLCmd.Flags().StringVarP(&flexOut, "output", "o", "table", "table|json|csv")

	flexSQLCmd.RegisterFlagCompletionFunc("output", utils.CompleteOutputFormat)

	rootCmd.AddCommand(flexSQLCmd)
}
