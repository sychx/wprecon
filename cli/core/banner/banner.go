package banner

import (
	"time"

	"github.com/AngraTeam/wprecon/internal/printer"
)

func Banner() {
	printer.Println(BAR)
	printer.Print(ASCII)
	printer.Println(BAR)
}

func XBanner(URL string, randomAgent bool, timeSleep int) {
	Banner()

	printer.Done("Target:\t", URL)

	if randomAgent {
		printer.Done("Random Agent:\t", randomAgent)
	}

	if timeSleep != 0 {
		printer.Done("Sleep Requests:\t", timeSleep, "second(s)")
	}

	printer.Done("Started in:\t", time.Now().Format("Monday Jan 02 15:04:05 2006"), "\n")
}
