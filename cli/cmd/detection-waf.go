package cmd

import (
	"fmt"
	"strings"

	"github.com/blackcrw/wpsgo/pkg/printer"
	wpsfinger "github.com/blackcrw/wpsgo/tools/wordpress/fingerprint"
	"github.com/spf13/cobra"
)

func Detectionwaf(cmd *cobra.Command) {
	target, _ := cmd.Flags().GetString("url")
	detectionWaf, _ := cmd.Flags().GetBool("detection-waf")

	/* Start WAF detection */
	switch detectionWaf {
	case true:
		has, _ := wpsfinger.WAF(target)

		if has {
			printer.Warning("Do you wish to continue ?! [Y/n] :")

			var question string
			if fmt.Scan(&question); strings.ToLower(question) != "y" {
				printer.Fatal("Exiting...")
			}
		}
	}
	/* End WAF detection */
}
