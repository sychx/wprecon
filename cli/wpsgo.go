package cli

import (
	"os"

	"github.com/blackcrw/wpsgo/cli/cmd"
	"github.com/blackcrw/wpsgo/internal"
	"github.com/blackcrw/wpsgo/pkg/gohttp"
	"github.com/blackcrw/wpsgo/pkg/printer" // This is color lib
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wpsgo",
	Short: "Wordpress Scanner Go",
	Long:  `Wpsgo (Wordpress Scanner Go) is a scanner based on wpscan, only done in golang to get better performance!`,
	Run: func(ccmd *cobra.Command, args []string) {
		cmd.Wpcheck(ccmd)
		cmd.Detectionwaf(ccmd)
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

	rootCmd.PersistentFlags().StringP("url", "u", "", "Target URL (Ex: http(s)://google.com/) "+printer.Required())
	rootCmd.PersistentFlags().BoolP("detection-waf", "d", false, "Detection WAF")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")

	rootCmd.MarkPersistentFlagRequired("url")
}

func initBanner() {
	target, _ := rootCmd.Flags().GetString("url")
	isURL, err := gohttp.IsValidURL(target)

	switch isURL {
	case false:
		internal.Banner()
		printer.Fatal(err)
	case true:
		internal.SBanner(target)
	default:
		internal.Banner()
	}

}
