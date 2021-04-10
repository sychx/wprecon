package main

import (
	"fmt"
	"strings"

	"github.com/blackbinn/wprecon/cli"
	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/internal/pkg/gohttp"
	"github.com/blackbinn/wprecon/internal/pkg/printer"
	"github.com/blackbinn/wprecon/internal/pkg/text"
	"github.com/blackbinn/wprecon/tools/wordpress/enumerate"
)

func init() {
	database.Memory.SetString("Target", "https://share.america.gov/")
	database.Memory.SetString("HTTP wp-content", "wp-content")

	var response = gohttp.SimpleRequest("https://share.america.gov/")
	database.Memory.SetString("HTTP Index Raw", response.Raw)
}

func main() {
	cli.Execute()

	var constructorPlugins = enumerate.NewPlugins(database.Memory.GetString("HTTP Index Raw"), database.Memory.GetString("HTTP wp-content"))

	var ntpl = printer.NewTopLine("Loading Plugins...")

	for _, plugin := range constructorPlugins.Passive() {
		ntpl.Done("Plugin:", plugin[1], printer.Underline+printer.Yellow+"(Enumerate Passive)"+printer.Reset)
		printer.NewTopics("Location:", database.Memory.GetString("Target")+database.Memory.GetString("HTTP wp-content")+"/plugins/"+plugin[1]+"/").Done()

		if len(plugin) >= 3 {
			for version, confidence := range text.PercentageOfVersions(strings.Split(plugin[2], ",")) {
				printer.NewTopics("Version:", fmt.Sprint(version)).Done()
				printer.NewTopics("Confidence:", fmt.Sprintf("%d%%", confidence)).Prefix("  ").Warning()
			}
		} else {
			printer.NewTopics("Version:", printer.Underline+printer.Red+"Version Not Identify"+printer.Reset).Danger()
		}

		printer.NewTopics("Match(s):").Done()
		for _, match := range strings.Split(plugin[0], ",") {
			printer.NewTopics(fmt.Sprint(match)).Prefix("  ").Warning()
		}

		printer.Println()
	}

	fmt.Println("—————————————————————————————————————\n")

	for _, plugin := range constructorPlugins.Aggressive() {
		ntpl.Done("Plugin:", plugin[1], printer.Underline+printer.Yellow+"(Enumerate Aggressive)"+printer.Reset)
		printer.NewTopics("Location:", database.Memory.GetString("Target")+database.Memory.GetString("HTTP wp-content")+"/plugins/"+plugin[1]+"/").Done()

		if len(plugin) >= 3 {
			for version, confidence := range text.PercentageOfVersions(strings.Split(plugin[2], ",")) {
				printer.NewTopics("Version:", fmt.Sprint(version)).Done()
				printer.NewTopics("Confidence:", fmt.Sprintf("%d%%", confidence)).Prefix("  ").Warning()
			}
		} else {
			printer.NewTopics("Version:", printer.Underline+printer.Red+"Version Not Identify"+printer.Reset).Danger()
		}

		printer.NewTopics("Match(s):").Done()
		for _, match := range strings.Split(plugin[0], ",") {
			printer.NewTopics(fmt.Sprint(match)).Prefix("  ").Warning()
		}

		printer.Println()
	}

	printer.Println("Olá " + printer.Cyan + "Mundo" + printer.Reset)
	printer.Println("Olá " + printer.Green + "Mundo" + printer.Reset)
	printer.Println("Olá " + printer.Red + "Mundo" + printer.Reset)
	printer.Println("Olá " + printer.Magenta + "Mundo" + printer.Reset)
	printer.Println("Olá " + printer.Black + "Mundo" + printer.Reset)
	printer.Println("Olá " + printer.Blue + "Mundo" + printer.Reset)
}
