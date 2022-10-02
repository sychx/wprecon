/*
	!!! Nota !!!

	Criar um canal, para passar as informações, via goroutine.
	As informações de plugins/themas, deverão ser printadas em tempo real.
	Ele deve procurar todas as informações do plugin e quando terminar já printar logo em seguida.
	Isso trará uma ideia de velocidade tendo em vista que tudo será printado na hora.

	// wp-includes/ms-settings.php","wp-includes/post-template.php",'wp-includes/shortcodes.php','wp-includes/rss-functions.php

*/

package main

import (
	"strings"

	"github.com/AngraTeam/wprecon/cli/core/banner"
	"github.com/AngraTeam/wprecon/cli/core/signal"
	database "github.com/AngraTeam/wprecon/internal/memory"
	"github.com/AngraTeam/wprecon/internal/printer"
	"github.com/spf13/cobra"
)

func init() {
	go signal.Exit()
	flags()
}

func main() {
	printer.HandlingFatal(root.Execute())
}

var root = &cobra.Command{
	Use:     "wprecon",
	Short:   "Wordpress Recon",
	Long:    `WPRecon, is a tool for the recognition of vulnerabilities and blackbox information for wordpress.`,
//	PreRun:  cmd.RootOptionsPreRun,
//	Run:     cmd.RootOptionsRun,
//	PostRun: cmd.RootOptionsPostRun,
}

func flags() {
	cobra.OnInitialize(options)

	root.PersistentFlags().StringP("url", "u", "", "Target URL (Ex: http(s)://example.com/). "+printer.REQUIRED)
	root.PersistentFlags().BoolP("verbose", "v", false, "Verbosity mode.")
	root.PersistentFlags().BoolP("tor", "", false, "Use Tor anonymity network")
	root.PersistentFlags().BoolP("disable-tls-checks", "", false, "Disables SSL/TLS certificate verification.")
	root.PersistentFlags().BoolP("random-agent", "", false, "Use randomly selected HTTP(S) User-Agent header value.")
	root.PersistentFlags().BoolP("force", "f", false, "Forces wprecon to not check if the target is running WordPress and forces other executions.")
	root.PersistentFlags().IntP("http-sleep", "", 0, "You can make each request slower, if there is a WAF, it can make it difficult for it to block you. (default: 0)")

	root.Flags().BoolP("aggressive-mode", "A", false, "Activates the aggressive mode of wprecon.")
	root.Flags().BoolP("detection-waf", "", false, "I will try to detect if the target is using any WAF Wordpress.")
	root.Flags().StringP("wp-content-dir", "", "wp-content", "In case the wp-content directory is customized.")

	root.MarkPersistentFlagRequired("url")

	root.SetHelpTemplate(banner.BANNER_HELP_ROOT)
}

func options() {
	var tor, _ = root.Flags().GetBool("tor")
	var force, _ = root.Flags().GetBool("force")
	var target, _ = root.Flags().GetString("url")
	var verbose, _ = root.Flags().GetBool("verbose")
	var httpSleep, _ = root.Flags().GetInt("http-sleep")
	var randomAgent, _ = root.Flags().GetBool("random-agent")
	var wpContentDir, _ = root.Flags().GetString("wp-content-dir")
	var disableTlsChecks, _ = root.Flags().GetBool("disable-tls-checks")

	if !strings.HasSuffix(target, "/") {
		database.Memory.SetString("Options URL", target+"/")
	} else {
		database.Memory.SetString("Options URL", target)
	}

	database.Memory.SetBool("Options Force", force)
	database.Memory.SetBool("HTTP Options TOR", tor)
	database.Memory.SetBool("Options Verbose", verbose)
	database.Memory.SetBool("HTTP Options Random Agent", randomAgent)
	database.Memory.SetBool("HTTP Options TLS Certificate Verify", disableTlsChecks)
	database.Memory.SetInt("HTTP Time Sleep", httpSleep)
	database.Memory.SetString("HTTP wp-content", wpContentDir)
}
