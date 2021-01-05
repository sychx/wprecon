package cmd

import (
	"os"

	"github.com/blackcrw/wprecon/cli/cmd"
	"github.com/blackcrw/wprecon/internal/pkg/banner"
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "wprecon",
	Short: "Wordpress Recon",
	Long:  `wprecon (Wordpress Recon) is a scanner based on wpscan, only done in golang to get better performance!`,
	Run:   cmd.RootOptionsRun,
}

var fuzzer = &cobra.Command{
	Use:     "fuzzer",
	Aliases: []string{"fuzz"},
	Short:   "Fuzzing directories or logins.",
	Long:    "This subcommand is mainly focused on fuzzing directories or logins.",
	Run:     cmd.FuzzerOptionsRun,
}

func ibanner() {
	verbose, _ := root.Flags().GetBool("verbose")
	target, _ := root.Flags().GetString("url")
	tor, _ := root.Flags().GetBool("tor")

	if isURL, err := gohttp.IsURL(target); isURL {
		banner.SBanner(target, tor, verbose)
	} else {
		banner.Banner()
		printer.Fatal(err)
	}
}

// Execute ::
func Execute() {
	if err := root.Execute(); err != nil {
		os.Exit(0)
	}
}

func init() {
	cobra.OnInitialize(ibanner)

	root.PersistentFlags().StringP("url", "u", "", "Target URL (Ex: http(s)://example.com/). "+printer.Required)
	root.PersistentFlags().BoolP("random-agent", "", false, "Use randomly selected HTTP(S) User-Agent header value.")
	root.PersistentFlags().BoolP("tor", "", false, "Use Tor anonymity network")
	root.PersistentFlags().BoolP("disable-tls-checks", "", false, "Disables SSL/TLS certificate verification.")
	root.PersistentFlags().BoolP("verbose", "v", false, "Verbosity mode.")
	root.PersistentFlags().BoolP("force", "f", false, "Forces wprecon to not check if the target is running WordPress.")

	root.Flags().BoolP("detection-waf", "", false, "I will try to detect if the target is using any WAF Wordpress.")
	root.Flags().BoolP("detection-honeypot", "", false, "I will try to detect if the target is a honeypot, based on the shodan.")
	root.Flags().BoolP("users-enumerate", "", false, "Use the supplied mode to enumerate Users.")
	root.Flags().BoolP("plugins-enumerate", "", false, "Use the supplied mode to enumerate Plugins.")
	root.Flags().BoolP("themes-enumerate", "", false, "Use the supplied mode to enumerate themes.")

	root.MarkPersistentFlagRequired("url")

	root.SetHelpTemplate(banner.HelpMain)

	// fuzzer
	fuzzer.Flags().StringP("wordlist", "w", "nil", "Wordlist")
	fuzzer.Flags().BoolP("backup-file", "b", false, "Performs a fuzzing to try to find the backup file if it exists.")
	fuzzer.Flags().StringP("usernames", "U", "", "Set usernames attack passwords.")
	fuzzer.Flags().StringP("passwords", "P", "", "Set usernames attack passwords.")
	fuzzer.Flags().StringP("attack-method", "M", "xml-rpc", "Avaliable: xml-rpc and wp-login")

	fuzzer.SetHelpTemplate(banner.HelpFuzzer)
	root.AddCommand(fuzzer)
}
