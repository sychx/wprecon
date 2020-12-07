package cli

import (
	"os"

	. "github.com/blkzy/wpsgo/cli/cmd"
	"github.com/blkzy/wpsgo/pkg/printer"
	color "github.com/logrusorgru/aurora" // This is color lib
	"github.com/spf13/cobra"
)

var requiredColorText = color.Red("(Required)").Bold().String()

var rootCmd = &cobra.Command{
	Use:          "wpsgo",
	SilenceUsage: true,
}

// Execute ::
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		printer.Danger(false, err)
		os.Exit(0)
	}
}

func init() {
	cobra.OnInitialize(cmd.initBanner)

	rootCmd.PersistentFlags().StringP("url", "u", "nil", "URL Target")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output (errors)")
}
