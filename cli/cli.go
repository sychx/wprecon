package cli

import (
	"github.com/AngraTeam/wprecon/cli/cmd"
	"github.com/AngraTeam/wprecon/cli/core/banner"
	"github.com/AngraTeam/wprecon/cli/core/signal"
	"github.com/AngraTeam/wprecon/internal/printer"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
    Use:     "wprecon",
    Short:   "Wordpress Recon",
    Long:    `WPRecon, is a tool for the recognition of vulnerabilities and blackbox information for wordpress.`,
    PreRun:  cmd.RootPreRun,
    Run:     cmd.RootRun,
    PostRun: cmd.RootPostRun,
}

func init() {
    var flag = root.Flags()

    flag.StringP("url", "u", "", "Target URL (Ex: http(s)://example.com/). " + printer.REQUIRED)
    flag.BoolP("verbose", "v", false, "Verbosity mode.")
    flag.Bool("tor", false, "User Tor anonymity network.")
    flag.Bool("disable-tls-checks", false, "Disables SSL/TLS certificate verification.")
    flag.Bool("random-agent", false, "Use randomly selected HTTP(S) User-Agent header value.")
    flag.Bool("detection-waf", false, "I will try to detect if the target is using any WAF Wordpress.")
    flag.BoolP("force", "f", false, "Forces wprecon to not check if the target is running WordPress and forces other executions.")
    flag.BoolP("aggressive-mode", "A", false, "Activates the aggressive mode of wprecon.")
    flag.Int("http-sleep", 0, "You can make each request slower, if there is a WAF, it can make it difficult for it to block you. (default: 0)")
    flag.String("wp-content-dir", "wp-content", "In case the wp-content directory is customized.")
    
    root.MarkFlagRequired("url")
    root.SetHelpTemplate(banner.HELP_ROOT)
}

func Start() {
    go signal.Exit()
    printer.HandlingFatal(root.Execute())
}