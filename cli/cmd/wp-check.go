package cmd

import (
	"fmt"
	"strings"

	"github.com/blackcrw/wprecon/pkg/printer"
	wpfinger "github.com/blackcrw/wprecon/tools/wordpress/fingerprint"
	"github.com/spf13/cobra"
)

func Wpcheck(cmd *cobra.Command) {
	target, _ := cmd.Flags().GetString("url")
	randomAgent, _ := cmd.Flags().GetBool("random-agent")
	nocheckwp, _ := cmd.Flags().GetBool("no-check-wp")

	switch nocheckwp {
	case false:
		hasWordpressValue := wpfinger.HasWordpress(target, randomAgent)
		hasWordpressValueString := fmt.Sprintf("%.2f%%", hasWordpressValue)

		if hasWordpressValue >= 62.5 {
			printer.Done("Wordpress confirmed with", hasWordpressValueString, "accuracy!")
		} else if hasWordpressValue < 62.5 && hasWordpressValue > 25.0 {
			printer.Warning("I'm not absolutely sure that this target is using wordpress!", hasWordpressValueString, "chance. do you wish to continue ? [Y/n]:")

			var question string
			if fmt.Scan(&question); strings.ToLower(question) != "y" {
				printer.Fatal("Exiting...")
			}
		} else {
			printer.Fatal("This target is not running wordpress!")
		}
	}
}
