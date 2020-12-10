package cmd

import (
	"github.com/blackcrw/wprecon/pkg/printer"
	wpscan "github.com/blackcrw/wprecon/tools/wordpress/scanner"
	"github.com/spf13/cobra"
)

func PluginEnum(cmd *cobra.Command) {
	target, _ := cmd.Flags().GetString("url")
	pluginenum, _ := cmd.Flags().GetBool("plugins-enum")
	randomUserAgent, _ := cmd.Flags().GetBool("random-agent")

	switch pluginenum {
	case true:
		lista := wpscan.PluginEnum(target, randomUserAgent)

		for _, name := range lista {
			if name != "" {
				printer.Done(name)
			}
		}

		if len(lista) == 0 {
			printer.Danger("No plugin was found...")
		}
	}

}

