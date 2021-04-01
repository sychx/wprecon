package cmd

import (
	"strings"

	"github.com/blackbinn/wprecon/internal/pkg/gohttp"
	"github.com/blackbinn/wprecon/internal/pkg/printer"
	"github.com/blackbinn/wprecon/internal/pkg/text"
	"github.com/blackbinn/wprecon/internal/pkg/wordlist"
	"github.com/blackbinn/wprecon/tools/wordpress/fuzzing"
	flag "github.com/spf13/cobra"
)

// Fuzzer :: This is the fuzzing flags variable
var Fuzzer = &flag.Command{
	Use:     "fuzzer",
	Aliases: []string{"fuzz"},
	Short:   "Fuzzing directories or logins.",
	Long:    "This subcommand is mainly focused on fuzzing directories or logins.",
	Run:     FuzzerOptionsRun,
	PostRun: FuzzerOptionsPostRun,
}

func FuzzerOptionsRun(cmd *flag.Command, args []string) {
	{
		var channel = make(chan *gohttp.Response)
		var math = len(wordlist.BackupFiles) * 3

		var ntpl = printer.NewTopLine("Initialize Fuzzer Backup")

		go fuzzing.BackupFile(channel)

		for i := 0; i <= math; i++ {
			select {
			case response := <-channel:

				switch {
				case response.Response.StatusCode == 200:
					ntpl.Done(printer.Bold+"URL:"+printer.Reset, response.URL.Full, printer.Bold+"Status Code:"+printer.Reset, response.Response.Status)
				case response.Response.StatusCode != 404:
					ntpl.Warning(printer.Bold+"URL:"+printer.Reset, response.URL.Full, printer.Bold+"Status Code:"+printer.Reset, response.Response.Status)
				case i == math:
					ntpl.Clean()
					ntpl.Endl()
				default:
					ntpl.Progress(math, printer.Bold+"URL:"+printer.Reset, response.URL.Full, printer.Bold+"Status Code:"+printer.Reset, response.Response.Status)
				}
			}
		}
	}

	{
		var channel = make(chan [3]string)
		var usernames = [3]string{"staff", "blk", "aaa"}

		var ntpl = printer.NewTopLine("Brute-Force to wp-login — Loading Wordlist... ")
		var passwords, count = text.ReadAllFile("/Users/blkz/MEGAsync/Pentest/Wordlists/namess.txt")

	loopXML:
		for _, username := range usernames {
			go fuzzing.XMLRPC(channel, username, passwords)

			for i := 1; i <= count; i++ {
				select {
				case response := <-channel:
					switch {
					case strings.Contains(strings.ToLower(response[0]), "admin"):
						ntpl.Done("Password Found!")
						printer.NewTopics("Username:", response[1]).Default()
						printer.NewTopics("Password:", response[2]).Default()
						printer.NewTopics("Method Used:", "wp-login").Default()
						printer.Println()
						printer.ResetSeek()

						continue loopXML
					case i >= count:
						ntpl.Danger("Password Not Found!")
						printer.NewTopics("Username:", response[1]).Default()
						printer.NewTopics("Method Used:", "wp-login").Default()
						printer.Println()
						printer.ResetSeek()

					default:
						ntpl.Progress(count, "Username:", response[1], "Passwords:", response[2])
					}
				}
			}
		}
	}

	{
		var channel = make(chan [3]string)
		var usernames = []string{"staff", "blk", "admin"}

		var ntpl = printer.NewTopLine("Brute-Force to wp-login — Loading Wordlist... ")
		var passwords, count = text.ReadAllFile("/Users/blkz/MEGAsync/Pentest/Wordlists/names.txt")

	loopWP:
		for _, username := range usernames {

			go fuzzing.WPLogin(channel, username, passwords)

			for i := 1; i <= count; i++ {
				select {
				case response := <-channel:
					switch {
					case response[0] == "302":
						ntpl.Done("Password Found!")
						printer.NewTopics("Username:", response[1]).Default()
						printer.NewTopics("Password:", response[2]).Default()
						printer.NewTopics("Method Used:", "wp-login").Default()
						printer.Println()
						printer.ResetSeek()

						continue loopWP

					case i >= count:
						ntpl.Danger("Password Not Found!")
						printer.NewTopics("Username:", response[1]).Default()
						printer.NewTopics("Method Used:", "wp-login").Default()
						printer.Println()
						printer.ResetSeek()

					default:
						ntpl.Progress(count, "Username:", response[1], "Passwords:", response[2])
					}
				}
			}
		}
	}
}
func FuzzerOptionsPostRun(cmd *flag.Command, args []string) {}
