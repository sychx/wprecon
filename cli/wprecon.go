package cli

import (
	"os"

	"github.com/blackcrw/wprecon/cli/cmd"
	"github.com/blackcrw/wprecon/internal"
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer" // This is color lib
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wprecon",
	Short: "Wordpress Recon",
	Long:  `wprecon (Wordpress Recon) is a scanner based on wpscan, only done in golang to get better performance!`,
	Run: func(ccmd *cobra.Command, args []string) {
		cmd.Wpcheck(ccmd)
		cmd.Detectionwaf(ccmd)
		cmd.UserEnum(ccmd)
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
	rootCmd.PersistentFlags().BoolP("detection-waf", "d", false, "I will try to detect if the target is using any WAF.")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbosity")
	rootCmd.PersistentFlags().BoolP("random-agent", "", false, "Use randomly selected HTTP(S) User-Agent header value")
	rootCmd.PersistentFlags().BoolP("user-enum", "e", false, "Tries to enumerate users")
	rootCmd.PersistentFlags().BoolP("no-check-wp", "", false, "Will skip wordpress check on target")

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
