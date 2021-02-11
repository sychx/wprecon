package cli

import (
	"os"
	"strings"

	"github.com/blackbinn/wprecon/cli/cmd"
	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/internal/pkg/banner"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:     "wprecon",
	Short:   "Wordpress Recon",
	Long:    `wprecon (Wordpress Recon) is a tool for wordpress exploration!`,
	Run:     cmd.RootOptionsRun,
	PostRun: cmd.RootOptionsPostRun,
}

var fuzzer = &cobra.Command{
	Use:     "fuzzer",
	Aliases: []string{"fuzz"},
	Short:   "Fuzzing directories or logins.",
	Long:    "This subcommand is mainly focused on fuzzing directories or logins.",
	Run:     cmd.FuzzerOptionsRun,
	PostRun: cmd.FuzzerOptionsPostRun,
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
	root.PersistentFlags().StringP("scripts", "", "", "Auxiliary scripts.")
	root.PersistentFlags().BoolP("random-agent", "", false, "Use randomly selected HTTP(S) User-Agent header value.")
	root.PersistentFlags().BoolP("tor", "", false, "Use Tor anonymity network")
	root.PersistentFlags().BoolP("disable-tls-checks", "", false, "Disables SSL/TLS certificate verification.")
	root.PersistentFlags().BoolP("verbose", "v", false, "Verbosity mode.")
	root.PersistentFlags().BoolP("force", "f", false, "Forces wprecon to not check if the target is running WordPress and forces other executions.")
	root.PersistentFlags().IntP("http-sleep", "", 0, "You can make each request slower, if there is a WAF, it can make it difficult for it to block you. (default: 0)")

	root.Flags().BoolP("aggressive-mode", "A", false, "Activates the aggressive mode of wprecon.")
	root.Flags().BoolP("detection-waf", "", false, "I will try to detect if the target is using any WAF Wordpress.")
	root.Flags().StringP("wp-content-dir", "", "wp-content", "In case the wp-content directory is customized.")

	root.MarkPersistentFlagRequired("url")

	root.SetHelpTemplate(banner.HelpMain)

	// fuzzer
	fuzzer.Flags().StringP("usernames", "U", "", "Set usernames attack passwords.")
	fuzzer.Flags().StringP("passwords", "P", "", "Set wordlist attack passwords.")
	fuzzer.Flags().BoolP("backup-file", "B", false, "Performs a fuzzing to try to find the backup file if it exists.")
	fuzzer.Flags().StringP("attack-method", "M", "xml-rpc", "Avaliable: xml-rpc and wp-login")

	fuzzer.SetHelpTemplate(banner.HelpFuzzer)
	root.AddCommand(fuzzer)
}

func ibanner() {
	if target, _ := root.Flags().GetString("url"); !strings.HasSuffix(target, "/") {
		Database.Target = target + "/"
	} else {
		Database.Target = target
	}

	Database.Force, _ = root.Flags().GetBool("force")
	Database.Verbose, _ = root.Flags().GetBool("verbose")
	Database.WPContent, _ = root.Flags().GetString("wp-content-dir")
	Database.OtherInformationsBool["http.options.tor"], _ = root.Flags().GetBool("tor")
	Database.OtherInformationsString["scripts.name"], _ = root.Flags().GetString("scripts")
	Database.OtherInformationsBool["http.options.randomuseragent"], _ = root.Flags().GetBool("random-agent")
	Database.OtherInformationsBool["http.options.tlscertificateverify"], _ = root.Flags().GetBool("tlscertificateverify")
	Database.OtherInformationsInt["http.requests.time.sleep"], _ = root.Flags().GetInt("http-sleep")

	if isURL := gohttp.IsURL(Database.Target); isURL {
		banner.SBanner()
	} else {
		banner.Banner()
	}

	func() {
		response := gohttp.SimpleRequest(Database.Target)

		Database.OtherInformationsString["target.http.index.raw"] = response.Raw
		Database.OtherInformationsString["target.http.index.server"] = response.Response.Header.Get("Server")
		Database.OtherInformationsString["target.http.index.php.version"] = response.Response.Header.Get("x-powered-by")
		Database.OtherInformationsString["target.http.index.cookie"] = response.Response.Header.Get("Set-Cookie")
	}()
}
