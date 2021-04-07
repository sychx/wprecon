package main

import (
	"fmt"
	"strings"

	"github.com/blackbinn/wprecon/cli"
	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/internal/pkg/gohttp"
	"github.com/blackbinn/wprecon/internal/pkg/printer"
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

	for i, plugin := range constructorPlugins.Passive() {
		printer.Done("Key:", fmt.Sprint(i), "Plugin:", plugin[1], "—", plugin[2], "—", plugin[0])
	}

	fmt.Println("—————————————————————————————————————\n")

	for i, plugin := range constructorPlugins.Aggressive() {
		printer.Done("Key:", fmt.Sprint(i), "Plugin:", plugin[1], "—", strings.Join(plugin[2:], ","), "—", plugin[0])
	}

	printer.Println("Olá " + printer.Cyan + "Mundo" + printer.Reset)
	printer.Println("Olá " + printer.Green + "Mundo" + printer.Reset)
	printer.Println("Olá " + printer.Red + "Mundo" + printer.Reset)
}
