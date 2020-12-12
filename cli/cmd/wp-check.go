package cmd

import (
	"fmt"
	"strings"

	. "github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	wpfinger "github.com/blackcrw/wprecon/tools/wordpress/fingerprint"
	"github.com/spf13/cobra"
)

func DetectionWP(cmd *cobra.Command) {
	target, _ := cmd.Flags().GetString("url")
	noCheckWp, _ := cmd.Flags().GetBool("no-check-wp")
	randomUserAgent, _ := cmd.Flags().GetBool("random-agent")
	tlsCertificateVerify, _ := cmd.Flags().GetBool("disable-tls-verify")

	switch noCheckWp {
	case false:
		var question string

		/* Why did I choose to pass a struct ?! Instead of direct values ​​for and etc. Simple! As each part can have a small change ... so I can have better control! And I avoid building several variables ... */
		optionsHttp := Http{
			URL:                  target,
			RandomUserAgent:      randomUserAgent,
			TLSCertificateVerify: tlsCertificateVerify}

		wordpressAccuracy := wpfinger.HasWordpress(optionsHttp)
		wordpressAccuracyString := fmt.Sprintf("%.2f%%", wordpressAccuracy)

		if wordpressAccuracy >= 62.5 {
			printer.Done("Wordpress confirmed with", wordpressAccuracyString, "accuracy!")
		} else if wordpressAccuracy < 62.5 && wordpressAccuracy > 25.0 {
			printer.Warning("I'm not absolutely sure that this target is using wordpress!", wordpressAccuracyString, "chance. do you wish to continue ? [Y/n]:")

			if fmt.Scan(&question); strings.ToLower(question) != "y" {
				printer.Fatal("Exiting...")
			}
		} else {
			printer.Fatal("This target is not running wordpress!")
		}
	}
}
