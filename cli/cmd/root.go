package cmd

import (
	flag "github.com/spf13/cobra"
)

// Root :: This is the global flags variable
var Root = &flag.Command{
	Use:     "wprecon",
	Short:   "Wordpress Recon",
	Long:    `wprecon (Wordpress Recon) is a tool for wordpress exploration!`,
	Run:     RootOptionsRun,
	PostRun: RootOptionsPostRun,
}

func RootOptionsRun(cmd *flag.Command, args []string)     {}
func RootOptionsPostRun(cmd *flag.Command, args []string) {}
