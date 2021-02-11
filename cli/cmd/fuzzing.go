package cmd

import (
	"fmt"
	"strings"

	"github.com/blackbinn/wprecon/pkg/printer"
	"github.com/blackbinn/wprecon/pkg/scripts"
	"github.com/blackbinn/wprecon/pkg/text"
	"github.com/blackbinn/wprecon/tools/wordpress/fuzzing"
	"github.com/spf13/cobra"

	. "github.com/blackbinn/wprecon/cli/config"
)

func FuzzerOptionsRun(cmd *cobra.Command, args []string) {
	backupfile, _ := cmd.Flags().GetBool("backup-file")
	attackmethod, _ := cmd.Flags().GetString("attack-method")

	usernames, _ := cmd.Flags().GetString("usernames")
	filePasswords, _ := cmd.Flags().GetString("passwords")

	if names := Database.OtherInformationsString["scripts.name"]; names != "" {
		for _, name := range strings.Split(names, ",") {
			printer.Done("Running Script:", name)

			s := scripts.NewScript()
			s.UseScript(name)
			s.Run()

			printer.Println()
		}
	}

	if backupfile {
		fuzzing.BackupFile()
		printer.Println()
	}

	if attackmethod == "xml-rpc" && usernames != "" || attackmethod == "xml-rpc" && filePasswords != "" {
		ntl := printer.NewTopLine(":: Brute-Force to XML-RPC — Loading wordlist... ::")

		passwords, _ := text.ReadAllFile(filePasswords)

		channel := make(chan [2]int)

		for _, username := range strings.Split(usernames, ",") {
			go fuzzing.XMLRPC(channel, username, passwords)

			for alive := true; alive; {
				select {
				case response := <-channel:
					var status = response[0]
					var password = passwords[response[1]]

					progress := ntl.Progress(len(passwords), "Username:", username, "Password:", password)

					if status == 1 {
						ntl.Done("I found the user password:", username)
						printer.List("Password:", password).Done()
						printer.List("Attack Method:", "XML-RPC").D().L()

						progress.Fill()
						alive = false
					} else if len(passwords) == 1+response[1] {
						ntl.Danger("No password worked for the user:", username)
						printer.Println()

						progress.Fill()
						alive = false
					}
				}
			}
		}
	}
	if attackmethod == "wp-login" {
		ntl := printer.NewTopLine(":: Brute-Force to wp-login — Loading Wordlist... ::")

		passwords, _ := text.ReadAllFile(filePasswords)

		channel := make(chan [2]int)

		for _, username := range strings.Split(usernames, ",") {
			go fuzzing.WPLogin(channel, username, passwords)

			for alive := true; alive; {
				select {
				case response := <-channel:
					var status = response[0]
					var password = passwords[response[1]]

					progress := ntl.Progress(len(passwords), "Username:", username, "Password:", password)

					if status == 1 {
						ntl.Done("I found the user password:", username)
						printer.List("Password:", password).Done()
						printer.List("Attack Method:", "WP-Login").D().L()

						progress.Fill()
						alive = false
					} else if len(passwords) == 1+response[1] {
						ntl.Danger("No password worked for the user:", username)
						printer.Println()

						progress.Fill()
						alive = false
					}
				}
			}
		}
	}

}

func FuzzerOptionsPostRun(cmd *cobra.Command, args []string) {
	printer.Done("Total requests:", fmt.Sprint(Database.TotalRequests))
}
