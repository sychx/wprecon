package cmd

import (
	"fmt"
	"strings"

	. "github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	wpfinger "github.com/blackcrw/wprecon/tools/wordpress/fingerprint"
	"github.com/spf13/cobra"
)

func DetectionWAF(cmd *cobra.Command) {
	target, _ := cmd.Flags().GetString("url")
	detectionWaf, _ := cmd.Flags().GetBool("detection-waf")
	randomUserAgent, _ := cmd.Flags().GetBool("random-agent")
	tlsCertificateVerify, _ := cmd.Flags().GetBool("disable-tls-verify")
	torProxy, _ := cmd.Flags().GetBool("tor")

	switch detectionWaf {
	case true:
		var question string

		printer.Loading("Active WAF detection module")

		/* Why did I choose to pass a struct ?! Instead of direct values ​​for and etc. Simple! As each part can have a small change ... so I can have better control! And I avoid building several variables ... */
		optionsHttp := Http{
			URL:                  target,
			RandomUserAgent:      randomUserAgent,
			Tor:                  torProxy,
			TLSCertificateVerify: tlsCertificateVerify}

		waf := wpfinger.WAF{
			Request: optionsHttp,
			Verbose: false}

		if has, status, name := waf.Detection(); has {
			printer.LoadingWarning(status, "—", "WAF :", name)

			printer.Warning("Do you wish to continue ?! [Y/n] :")
			if fmt.Scan(&question); strings.ToLower(question) != "y" {
				printer.Fatal("Exiting...")
			}
		} else {
			printer.LoadingDanger("No WAF was detected! But that doesn't mean it doesn't.")
		}
	}
}
