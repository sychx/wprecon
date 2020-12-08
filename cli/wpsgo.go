package cli

import (
	"fmt"
	"os"

	"github.com/blackcrw/wpsgo/internal"
	"github.com/blackcrw/wpsgo/pkg/gohttp"
	"github.com/blackcrw/wpsgo/pkg/printer" // This is color lib
	wpsfinger "github.com/blackcrw/wpsgo/tools/wordpress/fingerprint"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wpsgo",
	Short: "Wordpress Scanner Go",
	Long:  `Wpsgo (Wordpress Scanner Go) is a scanner based on wpscan, only done in golang to get better performance!`,
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("url")
		detectionWaf, _ := cmd.Flags().GetBool("detection-waf")

		hasWordpressValue := wpsfinger.HasWordpress(target)
		hasWordpressValueString := fmt.Sprintf("%.2f%%", hasWordpressValue)

		if hasWordpressValue >= 62.5 {
			printer.Done("Wordpress confirmed with", hasWordpressValueString, "accuracy!")
		} else if hasWordpressValue < 62.5 && hasWordpressValue > 25.0 {
			printer.Warning("I'm not absolutely sure that this target is using wordpress!", hasWordpressValueString, "chance. do you wish to continue ? [Y/n]:")

			var question string
			if fmt.Scan(&question); question != "Y" {
				printer.Fatal("Exiting...")
			}
		} else {
			printer.Fatal("This target is not running wordpress!")
		}

		switch detectionWaf {
		case true:
			waf, WafName := wpsfinger.WAF(target)

			if waf {
				printer.Warning("Yes! This is using a WAF. Name WAF:", WafName)
			} else {
				printer.Warning("Not! This is not using a WAF.")
			}
		}

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
