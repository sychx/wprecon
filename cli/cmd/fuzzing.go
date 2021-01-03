package cmd

import (
	"fmt"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/tools/wordpress/fuzzing"
	"github.com/spf13/cobra"
)

// FuzzerOptionsRun ::
func FuzzerOptionsRun(cmd *cobra.Command, args []string) {
	target, _ := cmd.Flags().GetString("url")
	tor, _ := cmd.Flags().GetBool("tor")
	verbose, _ := cmd.Flags().GetBool("verbose")
	tlscertificateverify, _ := cmd.Flags().GetBool("disable-tls-verify")
	randomuseragent, _ := cmd.Flags().GetBool("random-agent")

	backupfile, _ := cmd.Flags().GetBool("backup-file")

	options := &gohttp.HTTPOptions{
		URL: gohttp.URLOptions{
			Simple: target,
		},
		Options: gohttp.Options{
			Tor:                  tor,
			RandomUserAgent:      randomuseragent,
			TLSCertificateVerify: tlscertificateverify,
		},
	}

	if backupfile {
		FB := fuzzing.BackupDir{
			HTTP:    options,
			Verbose: verbose,
		}

		FB.Run()
	}

	printer.Done("Total requests:", fmt.Sprint(options.TotalRequests))
}
