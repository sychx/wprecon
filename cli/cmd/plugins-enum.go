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

		/* Why did I choose to pass a struct ?! Instead of direct values ​​for and etc. Simple! As each part can have a small change ... so I can have better control! And I avoid building several variables ... */
		optionsHttp := Http{
			URL:                  target,
			RandomUserAgent:      randomUserAgent,
			TLSCertificateVerify: tlsCertificateVerify}

		plugins := wpscan.Plugins{
			Request: optionsHttp,
			Verbose: false}

		if has, names := plugins.Enumerate(); has {
			for _, name := range names {
				if name != "" {
					printer.Done("Plugin:", name)

					if changelog, response := plugins.Changelog(name); changelog {
						printer.Warning("Changelog:", response.URLFULL)
					}
				}
			}
		} else {
			printer.Danger("No plugin was found!")
		}

	}

}
