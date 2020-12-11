package cmd

import (
	. "github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	wpscan "github.com/blackcrw/wprecon/tools/wordpress/scanner"
	"github.com/spf13/cobra"
)

func UsersEnum(cmd *cobra.Command) {
	target, _ := cmd.Flags().GetString("url")
	userEnum, _ := cmd.Flags().GetBool("users-enumerate")
	randomUserAgent, _ := cmd.Flags().GetBool("random-agent")
	tlsCertificateVerify, _ := cmd.Flags().GetBool("disable-tls-verify")

	switch userEnum {
	case true:
		printer.Loading("Hunting users...")

		optionsHttp := Http{
			URL:                  target,
			RandomUserAgent:      randomUserAgent,
			TLSCertificateVerify: tlsCertificateVerify}

		if has, users := wpscan.UserEnumJson(optionsHttp); has {
			for _, user := range users {
				printer.Done("User:", user.Name, "—", "Slug:", user.Slug)
			}
		} else if has, users := wpscan.UserEnumRss(optionsHttp); has {
			for _, user := range users {
				if user.Name != "" {
					printer.Done("User:", user.Name)
				}
			}
		} else {
			printer.LoadingDanger("Unfortunately no user was found. ;-;")
		}
	}

}
