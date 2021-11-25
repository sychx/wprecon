package main

import (
	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/printer"
)

func init() {
	database.Memory.SetString("Target", "https://www.techo.org/")
	database.Memory.SetString("HTTP wp-content", "wp-content")

	var response = net.SimpleRequest("https://www.techo.org/")
	database.Memory.SetString("HTTP Index Raw", response.Raw)
}

func main() {
	printer.Println("Hello " + printer.Cyan + "World" + printer.Reset)
	printer.Println("Hello " + printer.Green + "World" + printer.Reset)
	printer.Println("Hello " + printer.Red + "World" + printer.Reset)
	printer.Println("Hello " + printer.Magenta + "World" + printer.Reset)
	printer.Println("Hello " + printer.Black + "World" + printer.Reset)
	printer.Println("Hello " + printer.Blue + "World" + printer.Reset)
}