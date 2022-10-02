package signal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/AngraTeam/wprecon/internal/printer"
)

func Exit() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)

	<-sc

	printer.Printf("Your press CTRL+C\r\n")
	syscall.Exit(0)
}
