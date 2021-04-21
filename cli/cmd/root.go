package cmd

import (
	"fmt"
	"strings"

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

		tagEnumeratePassive    = printer.Underline + printer.Yellow + "(Enumerate Passive)" + printer.Reset
		tagEnumerateAggressive = printer.Underline + printer.Yellow + "(Enumerate Aggressive)" + printer.Reset
		tagNoVersion           = printer.Underline + printer.Red + "Version Not Identify" + printer.Reset

		channel = make(chan []string)

		plugins = enumerate.NewPlugins(target, indexraw, wpcontent)
		ntpl    = printer.NewTopLine("Loading Plugins...")
	)

	for _, plugin := range plugins.Passive() {
		var (
			name     = plugin[1]
			versions = plugin[2]
			url      = target + wpcontent + "/plugins/" + name + "/"
		)

		ntpl.Done("Plugin:", name, tagEnumeratePassive)
		printer.NewTopics("Location:", url).Done()

		for version, confidence := range text.PercentageOfVersions(strings.Split(versions, ",")) {
			printer_version(confidence, tagNoVersion, version)
			printer_match(strings.Split(plugin[0], ","), version)
		}

		printer.Println()
	}

	go plugins.Aggressive(channel)

	for done := true; done; {
		select {
		case plugin, ok := <-channel:
			if ok {
				ntpl.Done("Plugin:", plugin[1], tagEnumerateAggressive)
				printer.NewTopics("Location:", target+wpcontent+"/plugins/"+plugin[1]+"/").Done()

				for version, confidence := range text.PercentageOfVersions(strings.Split(plugin[2], ",")) {
					printer_version(confidence, tagNoVersion, version)
					printer_match(strings.Split(plugin[0], ","), version)
				}

				printer.Println()
			} else {
				done = false
			}
		}
	}

	if plugins.LenPluginsAggressive == 0 && plugins.LenPluginsPassive == 0 {
		ntpl.Danger("No Plugin Detected")
	}
}

func RootOptionsPostRun(cmd *flag.Command, args []string) {}
