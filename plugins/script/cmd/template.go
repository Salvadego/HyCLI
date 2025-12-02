package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
}

var (
	templateName string
	templateFile string
	scriptDir    string = os.Getenv("HYCLI_SCRIPT_HOME")
)

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.Flags().StringVarP(&templateFile, "file", "f", "", "File containing the script to be saved")
}

var templateCmd = &cobra.Command{
	Use:   "template [TEMPLATE_NAME]",
	Short: "Create a new script template",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if scriptDir == "" {
			return fmt.Errorf("Script dir couldn't be found.")
		}

		templateName = args[0]
		if templateFile == "" {
			return fmt.Errorf("No template file provided")
		}

		file_ext := strings.ToLower(filepath.Ext(templateFile))
		templateFileSave := fmt.Sprintf("%s/%s%s", scriptDir, templateName, file_ext)

		b, err := os.ReadFile(templateFile)
		if err != nil {
			return err
		}

		err = os.WriteFile(templateFileSave, b, 0644)

		return nil
	},
}
