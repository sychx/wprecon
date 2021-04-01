package main

import (
	"github.com/blackbinn/wprecon/cli"
	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/internal/pkg/printer"
)

func init() {
	database.Memory.SetString("Target", "http://192.168.99.100:8000/")
	database.Memory.SetString("HTTP wp-content", "wp-content/")
}

func main() {
	cli.Execute()

	printer.Println("Olá " + printer.Cyan + "Mundo" + printer.Reset)
	printer.Println("Olá " + printer.Green + "Mundo" + printer.Reset)
	printer.Println("Olá " + printer.Red + "Mundo" + printer.Reset)
}
