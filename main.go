package main

import (
	"github.com/blackbinn/wprecon/cli"
	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/internal/pkg/gohttp"
	"github.com/blackbinn/wprecon/internal/pkg/printer"
)

func init() {
	database.Memory.SetString("Target", "https://www.techo.org/")
	database.Memory.SetString("HTTP wp-content", "wp-content")

	var response = gohttp.SimpleRequest("https://www.techo.org/")
	database.Memory.SetString("HTTP Index Raw", response.Raw)
}

func main() {
	cli.Execute()

	printer.Println("Hello " + printer.Cyan + "World" + printer.Reset)
	printer.Println("Hello " + printer.Green + "World" + printer.Reset)
	printer.Println("Hello " + printer.Red + "World" + printer.Reset)
	printer.Println("Hello " + printer.Magenta + "World" + printer.Reset)
	printer.Println("Hello " + printer.Black + "World" + printer.Reset)
	printer.Println("Hello " + printer.Blue + "World" + printer.Reset)
}
