package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/internal/pkg/printer"
	"github.com/blackbinn/wprecon/internal/pkg/text"
	"github.com/blackbinn/wprecon/tools/wordpress/enumerate"
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

var (
	printer_version = func(confidence int, tag, version string) { var str string; if confidence == 0 { str = tag } else { str = fmt.Sprintf("%s (confidence: %d%%)", version, confidence) }; printer.NewTopics("Version:", str).Done() }
	printer_match = func(submatch []string, version string) { for _, match := range submatch { if strings.Contains(match, version) { printer.NewTopics(match).Prefix("  ").Warning() } } }
)

func RootOptionsRun(cmd *flag.Command, args []string) {
	var (
		target    string = database.Memory.GetString("Target")
		indexraw  string = database.Memory.GetString("HTTP Index Raw")
		wpcontent string = database.Memory.GetString("HTTP wp-content")

		tagNoVersion = printer.Underline + printer.Red + "Version Not Identify" + printer.Reset

		plugins = enumerate.NewPlugins(target, indexraw, wpcontent)
		themes  = enumerate.NewThemes(target, indexraw, wpcontent)

		ntpl = printer.NewTopLine("Loading Plugins...")

		aggressive_mode,_ = cmd.Flags().GetBool("aggressive-mode")
	)

	var channel = make(chan [5]string)

	if aggressive_mode {
		go plugins.Aggressive(channel)
	}else{
		go plugins.Passive(channel)
	}

	time.Sleep(2*time.Second)

	for i := 1; i <= plugins.LenPluginsAggressive+plugins.LenPluginsPassive; i++ {
		select {
		case informations := <-channel:
			var (
				name     = informations[1]
				versions = informations[2]
				matchs   = informations[0]
				tag      = printer.Underline + printer.Yellow + "("+informations[3]+")" + printer.Reset
				url      = target + wpcontent + "/plugins/" + name + "/"
			)
	
			ntpl.Done("Plugin:", name, tag)
			printer.NewTopics("Location:", url).Done()

			for version, confidence := range text.PercentageOfVersions(strings.Split(versions, ",")) {
				printer_version(confidence, tagNoVersion, version)
				printer_match(strings.Split(matchs, ","), version)
			}

			printer.Println()
		}
	}

	if aggressive_mode {
		go themes.Aggressive(channel)
	}else{
		go themes.Passive(channel)
	}

	time.Sleep(2*time.Second)

	for i := 1; i <= themes.LenThemesAggressive+themes.LenThemesPassive; i++ {
		select {
		case informations := <-channel:
			var (
				name     = informations[1]
				versions = informations[2]
				matchs   = informations[0]
				tag      = printer.Underline + printer.Yellow + "("+informations[3]+")" + printer.Reset
				url      = target + wpcontent + "/themes/" + name + "/"
			)
	
			ntpl.Done("Theme:", name, tag)
			printer.NewTopics("Location:", url).Done()

			for version, confidence := range text.PercentageOfVersions(strings.Split(versions, ",")) {
				printer_version(confidence, tagNoVersion, version)
				printer_match(strings.Split(matchs, ","), version)
			}

			printer.Println()
		}
	}


}

func RootOptionsPostRun(cmd *flag.Command, args []string) {}
