package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/AngraTeam/wprecon/cli/core/banner"
	"github.com/AngraTeam/wprecon/cli/core/verbose"
	"github.com/AngraTeam/wprecon/internal/http"
	database "github.com/AngraTeam/wprecon/internal/memory"
	"github.com/AngraTeam/wprecon/internal/printer"
	"github.com/AngraTeam/wprecon/tools/finders/plugins"
	"github.com/AngraTeam/wprecon/tools/interesting"
	"github.com/spf13/cobra"
)

func RootPreRun(cmd *cobra.Command, args []string) {
    var flag = cmd.Flags()

    var tor, _ = flag.GetBool("tor")
    var force, _ = flag.GetBool("force")
    var target, _ = flag.GetString("url")
    var verbose, _ = flag.GetBool("verbose")
    var httpSleep, _ = flag.GetInt("http-sleep")
    var randomAgent, _ = flag.GetBool("random-agent")
    var wpContentDir, _ = flag.GetString("wp-content-dir")
    var disableTlsChecks, _ = flag.GetBool("disable-tls-checks")

    if !strings.HasSuffix(target, "/") {
        database.Memory.SetString("Target URL", target+"/")
    } else {
        database.Memory.SetString("Target URL", target)
    }

    database.Memory.SetBool("Force", force)
    database.Memory.SetBool("HTTP TOR", tor)
    database.Memory.SetBool("Verbose", verbose)
    database.Memory.SetBool("HTTP Random Agent", randomAgent)
    database.Memory.SetBool("HTTP TLS Certificate Verify", disableTlsChecks)
    database.Memory.SetInt("HTTP Time Sleep", httpSleep)
    database.Memory.SetString("HTTP wp-content", wpContentDir)

    if err := http.ThisIsURL(target); err == nil {
        banner.XBanner(database.Memory.GetString("Target URL"), randomAgent, httpSleep)
    } else {
        banner.Banner()
        printer.HandlingFatal(err)
    }

    var response, err = http.Request(database.Memory.GetString("Target URL"))

    printer.HandlingFatal(err)

    database.Memory.SetString("HTTP Index Raw", response.Raw)
    database.Memory.SetString("HTTP PHP Version", response.Response.Header.Get("x-powered-by"))
    database.Memory.SetString("HTTP Server", response.Response.Header.Get("Server"))
    database.Memory.SetString("HTTP Index Cookie", response.Response.Header.Get("Set-Cookie"))
}

func RootRun(cmd *cobra.Command, args []string) {
    var target = database.Memory.GetString("Target URL")
    var index = database.Memory.GetString("HTTP Index Raw")

    if !database.Memory.GetBool("Force") {
        go verbose.Println("Checking if the site is a wordpress!")

        if confidence := interesting.WordpressCheck(target, index); confidence >= 40.0 {
            printer.Done("WordPress confirmed with", fmt.Sprint(confidence)+"%%", "confidence!\n")
        } else if confidence < 40.0 && confidence > 15.0 {
            if q := printer.ScanQ("I'm not absolutely sure that this target is using wordpress!", fmt.Sprint(confidence)+"%%", "chance. do you wish to continue ? [Y]es | [n]o : "); q != "y" && q != "\n" {
                printer.Fatal("Exiting...")
            }
        } else if confidence < 15.0 {
            printer.Fatal("This target is not running wordpress!")
        }
    }

    go verbose.Println("Searching the core wordpress version!")
    if version := interesting.WordPressVersion(index); version != "" {
        printer.Done("WordPress Version:", version)
    } else { printer.Danger("I couldn't find the core wordpress version") }

    printer.Println()

    verbose.Println("Passively looking for plugins and their versions.")
    plugins.PassiveVersionSearch()

    defer printer.Println()
}

func RootPostRun(cmd *cobra.Command, args []string) {
    printer.Info("Total requests:", database.Memory.GetInt("HTTP Total"))
    printer.Info("Ended in:", time.Now().Format("02 Jan 15:04:05"))
}
