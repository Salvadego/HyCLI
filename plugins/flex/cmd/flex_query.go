package cmd

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flex/internal"
	"flex/utils"
	"fmt"
	"os"

	"github.com/Salvadego/hac/hac"
	"github.com/spf13/cobra"
)

var (
	flexQueryFile string
	flexOut       string
	flexQueryMax  int
	flexQuerySQL  bool
)

var flexQueryCmd = &cobra.Command{
	Use:   "query [FLEXSEARCH_QUERY]",
	Short: "Run a FlexibleSearch query",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var q string
		if flexQueryFile != "" {
			b, err := os.ReadFile(flexQueryFile)
			if err != nil {
				return err
			}
			q = string(b)
		} else {
			if len(args) == 0 {
				fmt.Println("FlexibleSearch query is required unless --file is used")
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
			FlexibleSearchQuery: q,
			User:                username,
			MaxCount:            flexQueryMax,
		}, nil)
		if err != nil {
			return err
		}

		if flexQuerySQL {
			fmt.Println(resp.Query)
			return nil
		}

		flexQueryOutput(resp, OutputType(flexOut))

		return nil
	},
}

func flexQueryOutput(resp *hac.FlexSearchResponse, out OutputType) error {
	switch out {
	case JSONOutput:
		b, err := json.MarshalIndent(resp.ResultList, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil

	case CSVOutput:
		w := csv.NewWriter(os.Stdout)
		if err := w.Write(resp.Headers); err != nil {
			return err
		}
		for _, row := range resp.ResultList {
			strRow := make([]string, len(row))
			for i, v := range row {
				strRow[i] = fmt.Sprint(v)
			}
			if err := w.Write(strRow); err != nil {
				return err
			}
		}
		w.Flush()
		return w.Error()

	default:
		fmt.Println("Headers:", resp.Headers)
		for _, row := range resp.ResultList {
			fmt.Println(row)
		}
		return nil
	}
}

func init() {
	flexQueryCmd.Flags().StringVarP(&flexOut, "output", "o", "table", "table|json|csv")
	flexQueryCmd.RegisterFlagCompletionFunc("output", utils.CompleteOutputFormat)

	flexQueryCmd.Flags().StringVarP(&flexQueryFile, "file", "f", "", "File with FlexibleSearch query")
	flexQueryCmd.Flags().IntVarP(&flexQueryMax, "max", "m", 100, "Max results")
	flexQueryCmd.Flags().BoolVar(&flexQuerySQL, "sql", false, "Print only generated SQL")

	rootCmd.AddCommand(flexQueryCmd)
}
