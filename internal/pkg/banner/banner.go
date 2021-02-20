package banner

import (
	"fmt"
	"strings"
	"time"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/internal/pkg/update"
	"github.com/blackbinn/wprecon/internal/pkg/version"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
	"github.com/blackbinn/wprecon/pkg/scripts"
)

// Banner :: A simple banner.
func Banner() {
	printer.Println("——————————————————————————————————————————————————————————————————")
	printer.Println("___       ______________________________________________   __\n__ |     / /__  __ \__  __ \__  ____/_  ____/_  __ \__  | / /\n__ | /| / /__  /_/ /_  /_/ /_  __/  _  /    _  / / /_   |/ / \n__ |/ |/ / _  ____/_  _, _/_  /___  / /___  / /_/ /_  /|  /  \n____/|__/  /_/     /_/ |_| /_____/  \____/  \____/ /_/ |_/   \n")
	printer.Println("Github: ", "https://github.com/blackbinn/wprecon")

	if newVersion := update.CheckUpdate(); newVersion != "" {
		printer.Println("Version:", version.Version, "(New Version: "+newVersion+")")
	} else {
		printer.Println("Version:", version.Version)
	}

	printer.Println("——————————————————————————————————————————————————————————————————")
}

// SBanner :: A banner that will only be executed if the scan is started correctly.
func SBanner() {
	Banner()

	printer.Done("Target:     ", Database.Target)

	if Database.OtherInformationsBool["http.options.tor"] {
		ipTor := gohttp.TorGetIP()

		printer.Done("Proxy:      ", ipTor)
	}

	Database.TimeStart = time.Now().Format(("Monday Jan 02 15:04:05 2006"))

	printer.Done("Started in: ", Database.TimeStart)

	if names := Database.OtherInformationsString["scripts.name"]; names != "" {
		var names = strings.Split(names, ",")

		for _, name := range names {
			if !scripts.Exists(name) {
				printer.Fatal("The \"" + name + "\" script does not exist")
			}
		}

		printer.Done("Loaded:     ", fmt.Sprintf("%d", len(names)), "Script(s)...")
	}

	if Database.Verbose && Database.OtherInformationsBool["http.options.tor"] {
		printer.Danger("(Alert) Activating verbose mode together with tor mode can make the wprecon super slow. \n")
	} else if Database.Verbose {
		printer.Danger("(Alert) Enabling verbose mode can slow down wprecon. \n")
	} else {
		printer.Println()
	}
}
