package cli

import (
	"os"
	"strings"

	"github.com/blackbinn/wprecon/cli/cmd"
	"github.com/blackbinn/wprecon/internal/database"
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
	fuzzer.Flags().StringP("p-prefix", "", "", "Sets a prefix for all passwords in the wordlist.")
	fuzzer.Flags().StringP("p-suffix", "", "", "Sets a suffix for all passwords in the wordlist.")

	fuzzer.SetHelpTemplate(banner.HelpFuzzer)
	root.AddCommand(fuzzer)
}

func ibanner() {
	if target, _ := root.Flags().GetString("url"); !strings.HasSuffix(target, "/") {
		database.Memory.SetString("Target", target+"/")
	} else {
		database.Memory.SetString("Target", target)
	}
	x1, _ := root.Flags().GetBool("force")
	database.Memory.SetBool("Force", x1)
	x2, _ := root.Flags().GetBool("tor")
	database.Memory.SetBool("HTTP Options TOR", x2)
	x3, _ := root.Flags().GetBool("verbose")
	database.Memory.SetBool("Verbose", x3)
	x4, _ := root.Flags().GetString("wp-content-dir")
	database.Memory.SetString("HTTP wp-content", x4)
	x5, _ := root.Flags().GetString("scripts")
	database.Memory.SetString("Scripts List Names", x5)
	x6, _ := fuzzer.Flags().GetString("p-prefix")
	database.Memory.SetString("Passwords Prefix", x6)
	x7, _ := fuzzer.Flags().GetString("p-suffix")
	database.Memory.SetString("Passwords Suffix", x7)
	x8, _ := root.Flags().GetBool("random-agent")
	database.Memory.SetBool("HTTP Options Random Agent", x8)
	x9, _ := root.Flags().GetBool("tlscertificateverify")
	database.Memory.SetBool("HTTP Options TLS Certificate Verify", x9)
	x10, _ := root.Flags().GetInt("http-sleep")
	database.Memory.SetInt("HTTP Time Sleep", x10)

	if isURL := gohttp.IsURL(database.Memory.GetString("Target")); isURL {
		banner.SBanner()
	} else {
		banner.Banner()
	}

	func() {
		response := gohttp.SimpleRequest(database.Memory.GetString("Target"))

		database.Memory.SetString("HTTP Index Raw", response.Raw)
		database.Memory.SetString("HTTP PHP Version", response.Response.Header.Get("x-powered-by"))
		database.Memory.SetString("HTTP Server", response.Response.Header.Get("Server"))
		database.Memory.SetString("HTTP Index Cookie", response.Response.Header.Get("Set-Cookie"))
	}()
}
