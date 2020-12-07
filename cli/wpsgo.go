package cli

import (
	"os"

	"github.com/blkzy/wpsgo/internal"
	"github.com/blkzy/wpsgo/pkg/gohttp"
	"github.com/blkzy/wpsgo/pkg/printer" // This is color lib
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wpsgo",
	Short: "Wordpress Scanner Go",
	Long:  `Wpsgo (Wordpress Scanner Go) is a scanner based on wpscan, only done in golang to get better performance!`,
	Run: func(cmd *cobra.Command, args []string) {
		printer.Danger("IALA")
	},
}

// Execute ::
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// printer.Danger(err)
		os.Exit(0)
	}
}

func init() {
	cobra.OnInitialize(initBanner)

	rootCmd.PersistentFlags().StringP("url", "u", "", "URL Target "+printer.Required())
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")

	rootCmd.MarkPersistentFlagRequired("url")
}

func initBanner() {
	target, _ := rootCmd.Flags().GetString("url")

	x, _ := gohttp.IsValidURL(target)

	printer.Println(x)

	if target != "" {
		internal.SBanner(target)
	} else {
		internal.Banner()
	}
}
