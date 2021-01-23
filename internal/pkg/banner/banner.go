package banner

import (
	"time"

	. "github.com/blackcrw/wprecon/cli/config"
	"github.com/blackcrw/wprecon/internal/pkg/version"
	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/pkg/scripts"
)

// Banner :: A simple banner.
func Banner() {
	printer.Println("——————————————————————————————————————————————————————————————————")
	printer.Println(`___       ______________________________________________   __
__ |     / /__  __ \__  __ \__  ____/_  ____/_  __ \__  | / /
__ | /| / /__  /_/ /_  /_/ /_  __/  _  /    _  / / /_   |/ / 
__ |/ |/ / _  ____/_  _, _/_  /___  / /___  / /_/ /_  /|  /  
____/|__/  /_/     /_/ |_| /_____/  \____/  \____/ /_/ |_/   
`)
	printer.Println("Github: ", "https://github.com/blackcrw/wprecon")
	printer.Println("Version: ", version.Version)
	printer.Println("——————————————————————————————————————————————————————————————————")
}

// SBanner :: A banner that will only be executed if the scan is started correctly.
func SBanner() {
	Banner()

	printer.Done("Target:\t", InfosWprecon.Target)

	if InfosWprecon.OtherInformationsBool["http.options.tor"] {
		ipTor := gohttp.TorGetIP()

		printer.Done("Proxy:\t", ipTor)
	}

	InfosWprecon.TimeStart = time.Now().Format(("02/Jan/2006 15:04:05"))

	printer.Done("Starting:\t", InfosWprecon.TimeStart)

	if name := InfosWprecon.OtherInformationsString["scripts.name"]; name != "" {
		_, infos := scripts.Initialize(name)

		printer.Done("Script Name:", infos.Title)
		printer.Done("Script Desc:", infos.Description)
	}

	if InfosWprecon.Verbose && InfosWprecon.OtherInformationsBool["http.options.tor"] {
		printer.Danger("(Alert) Activating verbose mode together with tor mode can make the wprecon super slow. \n")
	} else if InfosWprecon.Verbose {
		printer.Danger("(Alert) Enabling verbose mode can slow down wprecon. \n")
	} else {
		printer.Println("")
	}
}
