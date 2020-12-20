package banner

import (
	"fmt"
	"time"

	version "github.com/blackcrw/wprecon/internal/pkg/version"
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
)

// Banner :: A simple banner.
func Banner() {
	fmt.Print("—————————————————————————————————————————————————————————————————————\n")
	fmt.Print(`
___       ______________________________________________   __
__ |     / /__  __ \__  __ \__  ____/_  ____/_  __ \__  | / /
__ | /| / /__  /_/ /_  /_/ /_  __/  _  /    _  / / /_   |/ / 
__ |/ |/ / _  ____/_  _, _/_  /___  / /___  / /_/ /_  /|  /  
____/|__/  /_/     /_/ |_| /_____/  \____/  \____/ /_/ |_/   

`)
	fmt.Print("Github: ", "https://github.com/blackcrw/wprecon\n")
	fmt.Print("Version: ", version.Version, "\n")
	fmt.Print("—————————————————————————————————————————————————————————————————————\n")
}

// SBanner :: A banner that will only be executed if the scan is started correctly.
func SBanner(target string, tor bool, verbose bool) {
	Banner()
	printer.Done("Target:", target)

	if tor {
		ipTor, err := gohttp.TorCheck()

		if err != nil {
			printer.Fatal(err)
		}

		printer.Done("Proxy:", ipTor)
	}

	timeFormat := time.Now().Format(("02/Jan/2006 15:04:05"))

	printer.Done("Starting:", timeFormat)

	if verbose && tor {
		printer.Danger("(Alert) Activating verbose mode together with tor mode can make the wprecon super slow.", "\n")
	} else if verbose {
		printer.Danger("(Alert) Enabling verbose mode can slow down wprecon.", "\n")
	} else {
		fmt.Println("")
	}
}
