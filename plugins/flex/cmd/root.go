package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	baseURL    string = os.Getenv("HYCLI_CLIENT_URL")
	username   string = os.Getenv("HYCLI_CLIENT_USER")
	password   string = os.Getenv("HYCLI_CLIENT_PASSWORD")
	skipVerify bool
)

var rootCmd = &cobra.Command{
	Use:   "flex",
	Short: "FlexibleSearch CLI for HAC",
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&skipVerify, "skip-verify", true, "Skip TLS verify")
}

func Execute() { rootCmd.Execute() }
