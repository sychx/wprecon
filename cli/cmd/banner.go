package cmd

import (
	"github.com/blkzy/wpsgo/internal"
	"github.com/spf13/cobra"
)

func InitBanner() {
	var cmd *cobra.Command

	target, _ := cmd.Flags().GetString("url")

	if target != "nil" {
		internal.SBanner(target)
	} else {
		internal.Banner()
	}
}
