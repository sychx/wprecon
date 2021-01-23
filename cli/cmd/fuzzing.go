package cmd

import (
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/pkg/scripts"
	"github.com/blackcrw/wprecon/tools/wordpress/fuzzing"
	"github.com/spf13/cobra"

	. "github.com/blackcrw/wprecon/cli/config"
)

func FuzzerOptionsRun(cmd *cobra.Command, args []string) {
	backupfile, _ := cmd.Flags().GetBool("backup-file")
	attackmethod, _ := cmd.Flags().GetString("attack-method")

	InfosWprecon.OtherInformationsString["target.http.fuzzing.usernames"], _ = cmd.Flags().GetString("usernames")
	InfosWprecon.OtherInformationsString["target.http.fuzzing.passwords.file.wordlist"], _ = cmd.Flags().GetString("passwords")

	if scriptsS := InfosWprecon.OtherInformationsString["scripts.name"]; scriptsS != "" {
		L, _ := scripts.Initialize(scriptsS)

		scripts.Run(L)
	}

	if backupfile {
		fuzzing.BackupFile()
		printer.Println()
	}

	if attackmethod == "xml-rpc" && InfosWprecon.OtherInformationsString["target.http.fuzzing.usernames"] != "" || attackmethod == "xml-rpc" && InfosWprecon.OtherInformationsString["target.http.fuzzing.passwords.file.wordlist"] != "" {
		printer.Done(":: Brute-Force to xml-rpc ::")
		fuzzing.XMLRPC()
	} else if attackmethod == "wp-login" && InfosWprecon.OtherInformationsString["target.http.fuzzing.usernames"] != "" || attackmethod == "wp-login" && InfosWprecon.OtherInformationsString["target.http.fuzzing.passwords.file.wordlist"] != "" {
		printer.Done(":: Brute-Force to wp-login ::")
		fuzzing.WPLogin()
	}
}
