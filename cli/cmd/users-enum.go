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
		printer.Warning("Hunting users...")

		/* Why did I choose to pass a struct ?! Instead of direct values ​​for and etc. Simple! As each part can have a small change ... so I can have better control! And I avoid building several variables ... */
		optionsHttp := Http{
			URL:                  target,
			RandomUserAgent:      randomUserAgent,
			TLSCertificateVerify: tlsCertificateVerify}

		users := wpscan.Users{
			Request: optionsHttp,
			Verbose: false}

		if has, usersJson := users.EnumerateJson(); has {
			for _, user := range usersJson {
				printer.Done("User:", user.Name)
			}
		} else if has, userRss := users.EnumerateRss(); has {
			for _, user := range userRss {
				if user.Name != "" {
					printer.Done("User:", user.Name)
				}
			}
		} else {
			printer.Danger("Unfortunately no user was found. ;-;")
		}
	}

}
