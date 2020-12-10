package cmd

import (
	wpscan "github.com/blackcrw/wprecon/tools/wordpress/scanner"
	"github.com/spf13/cobra"
)

func UserEnum(cmd *cobra.Command) {
	target, _ := cmd.Flags().GetString("url")
	userenum, _ := cmd.Flags().GetBool("user-enum")
	randomUserAgent, _ := cmd.Flags().GetBool("random-agent")

	switch userenum {
	case true:
		wpscan.UserEnum(target, randomUserAgent)
	}

}
