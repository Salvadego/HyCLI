package cmd

import (
	"context"
	"fmt"
	"hycli-log/internal"
	"hycli-log/utils"
	"strings"

	"github.com/Salvadego/hac/hac"
	"github.com/spf13/cobra"
)

var (
	loggerName string
	logLevel   string
)

func init() {
	changeLogCmd.Flags().StringVarP(&logLevel, "level", "l", "", "Log level")
	changeLogCmd.RegisterFlagCompletionFunc("level", utils.CompleteLogLevel)
	rootCmd.AddCommand(changeLogCmd)
}

var changeLogCmd = &cobra.Command{
	Use:               "change_log [logger_name]",
	Short:             "Changes the log level of a logger",
	Args:              cobra.ArbitraryArgs,
	ValidArgsFunction: CompleteLoggers,
	RunE:              RunChangeLogCmd,
}

func RunChangeLogCmd(cmd *cobra.Command, args []string) error {

	if len(args) == 0 {
		fmt.Println("Logger name is requeried")
		return cmd.Help()
	}

	if logLevel == "" {
		fmt.Println("Log level is requeried")
		return cmd.Help()
	}

	loggerName = args[0]

	c := internal.New(baseURL, username, password, skipVerify)
	ctx := context.Background()

	if err := internal.Login(c, ctx); err != nil {
		return err
	}

	level := hac.LogLevelName(logLevel)
	resp, err := c.Log.ChangeLogLevel(ctx, loggerName, level)
	if err != nil {
		return err
	}

	fmt.Printf("%s is now: %s\n", resp.LoggerName, resp.LevelName)

	return nil
}

func CompleteLoggers(
	cmd *cobra.Command,
	args []string,
	toComplete string,
) ([]string, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	ctx := context.Background()

	c := internal.New(baseURL, username, password, skipVerify)
	if err := internal.Login(c, ctx); err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	loggersResp, err := c.Log.GetCurrentLoggers(ctx)

	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	var completions []string
	for _, logger := range loggersResp {
		if strings.HasPrefix(logger.Name, toComplete) {
			item := fmt.Sprintf("%s\t(%s)", logger.Name, logger.EffectiveLevel.StandardLevel)
			completions = append(completions, item)
		}
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}
