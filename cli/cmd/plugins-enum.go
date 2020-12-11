package cmd

import (
	. "github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	wpscan "github.com/blackcrw/wprecon/tools/wordpress/scanner"
	"github.com/spf13/cobra"
)

func PluginsEnum(cmd *cobra.Command) {
	target, _ := cmd.Flags().GetString("url")
	pluginEnum, _ := cmd.Flags().GetBool("plugins-enumerate")
	randomUserAgent, _ := cmd.Flags().GetBool("random-agent")
	tlsCertificateVerify, _ := cmd.Flags().GetBool("disable-tls-verify")

	switch pluginEnum {
	case true:
		printer.Warning("Hunting plugins...")

		optionsHttp := Http{
			URL:                  target,
			RandomUserAgent:      randomUserAgent,
			TLSCertificateVerify: tlsCertificateVerify}

		if has, lista := wpscan.PluginsFind(optionsHttp); has {
			for _, name := range lista {
				if name != "" {
					printer.Done("Plugin:", name)
				}
			}
		} else {
			printer.Danger("No plugin was found!")
		}

	}

}
