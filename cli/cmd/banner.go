package cmd

import "github.com/blkzy/wpsgo/internal"

func initBanner() {
	target := cli.rootCmd.Flags().GetString("url")

	if target != "nil" {
		internal.SBanner()
	} else {
		internal.Banner()
	}
}
