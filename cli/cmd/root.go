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

func RootOptionsRun(cmd *flag.Command, args []string) {
	var channel = make(chan []string)
	var constructorPlugins = enumerate.NewPlugins(database.Memory.GetString("Target"), database.Memory.GetString("HTTP Index Raw"), database.Memory.GetString("HTTP wp-content"))
	var ntpl = printer.NewTopLine("Loading Plugins...")

	for _, plugin := range constructorPlugins.Passive() {
		ntpl.Done("Plugin:", plugin[1], printer.Underline+printer.Yellow+"(Enumerate Passive)"+printer.Reset)
		printer.NewTopics("Location:", database.Memory.GetString("Target")+database.Memory.GetString("HTTP wp-content")+"/plugins/"+plugin[1]+"/").Done()

		if len(plugin) >= 3 {
			for version, confidence := range text.PercentageOfVersions(strings.Split(plugin[2], ",")) {
				printer.NewTopics("Version:", fmt.Sprint(version)).Done()
				printer.NewTopics("Confidence:", fmt.Sprintf("%d%%", confidence)).Prefix("  ").Warning()
				for _, match := range strings.Split(plugin[0], ",") {
					if strings.Contains(match, version) {
						printer.NewTopics(fmt.Sprint(match)).Prefix("  ").Warning()
					}
				}
			}
		} else {
			printer.NewTopics("Version:", printer.Underline+printer.Red+"Version Not Identify"+printer.Reset).Danger()
		}

		printer.Println()
	}

	go constructorPlugins.Aggressive(channel)

	for done := true; done; {

		select {
		case plugin, ok := <-channel:
			if ok {
				ntpl.Done("Plugin:", plugin[1], printer.Underline+printer.Yellow+"(Enumerate Aggressive)"+printer.Reset)
				printer.NewTopics("Location:", database.Memory.GetString("Target")+database.Memory.GetString("HTTP wp-content")+"/plugins/"+plugin[1]+"/").Done()

				if len(plugin) == 3 {
					for version, confidence := range text.PercentageOfVersions(strings.Split(plugin[2], ",")) {
						printer.NewTopics("Version:", fmt.Sprint(version)).Done()
						printer.NewTopics("Confidence:", fmt.Sprintf("%d%%", confidence)).Prefix("  ").Warning()
						for _, match := range strings.Split(plugin[0], ",") {
							if strings.Contains(match, version) {
								printer.NewTopics(fmt.Sprint(match)).Prefix("  ").Warning()
							}
						}
					}
				} else {
					printer.NewTopics("Version:", printer.Underline+printer.Red+"Version Not Identify"+printer.Reset).Danger()
				}

				printer.Println()

			} else {
				done = false
			}
		}
	}

	if constructorPlugins.CountPluginsAggressive == 0 && constructorPlugins.CountPluginsPassive == 0 {
		ntpl.Danger("No Plugin Detected")
	}
}

func RootOptionsPostRun(cmd *flag.Command, args []string) {}
