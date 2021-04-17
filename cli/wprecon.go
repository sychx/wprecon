package cli

import (
	"os"

	. "github.com/blackbinn/wprecon/cli/cmd"
	"github.com/blackbinn/wprecon/internal/pkg/extensions"
	"github.com/blackbinn/wprecon/internal/pkg/printer"
)

// Execute :: This function is the one that will start the cli flags, and if there is an error in the cli it will automatically end the wprecon.
func Execute() {
	if err := Root.Execute(); err != nil {
		os.Exit(0)
	}
}

func init() {
	// Any flags that are set using the global variable will appear even when you want to use a subcommand.
	var global = Root.PersistentFlags()
	// All variables defined with the "local" variable will not appear in other subcommands.
	var local = Root.Flags()
	// All variables defined with the variable "fuzzer" will be available when using the subcommand "fuzzer or fuzz".
	var fuzzer = Fuzzer.Flags()

	global.StringP("url", "u", "", "Target URL (Ex: http(s)://example.com/). "+printer.Required)
	global.StringP("scripts", "", "", "Auxiliary scripts.")
	global.BoolP("random-agent", "", false, "Use randomly selected HTTP(S) User-Agent header value.")
	global.BoolP("tor", "", false, "Use Tor anonymity network")
	global.BoolP("disable-tls-checks", "", false, "Disables SSL/TLS certificate verification.")
	global.BoolP("verbose", "v", false, "Verbosity mode.")
	global.BoolP("force", "f", false, "Forces wprecon to not check if the target is running WordPress and forces other executions.")
	global.IntP("http-sleep", "", 0, "You can make each request slower, if there is a WAF, it can make it difficult for it to block you. (default: 0)")

	local.BoolP("aggressive-mode", "A", false, "Activates the aggressive mode of wprecon.")
	local.BoolP("detection-waf", "", false, "I will try to detect if the target is using any WAF Wordpress.")
	local.StringP("wp-content-dir", "", "wp-content", "In case the wp-content directory is customized.")

	fuzzer.StringP("usernames", "U", "", "Set usernames attack passwords.")
	fuzzer.StringP("passwords", "P", "", "Set wordlist attack passwords.")
	fuzzer.StringP("attack-method", "M", "xml-rpc", "Avaliable: xml-rpc and wp-login")
	fuzzer.StringP("p-prefix", "", "", "Sets a prefix for all passwords in the wordlist.")
	fuzzer.StringP("p-suffix", "", "", "Sets a suffix for all passwords in the wordlist.")
	fuzzer.BoolP("backup-file", "B", false, "Performs a fuzzing to try to find the backup file if it exists.")

	Fuzzer.SetHelpTemplate(extensions.BannerHelpFuzzer)
	Root.SetHelpTemplate(extensions.BannerHelpRoot)
	Root.MarkPersistentFlagRequired("url")
	Root.AddCommand(Fuzzer)
}
