package banner

import (
	"fmt"
	"strings"
	"time"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/internal/pkg/update"
	"github.com/blackbinn/wprecon/internal/pkg/version"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
	"github.com/blackbinn/wprecon/pkg/scripts"
)

// Banner :: A simple banner.
func Banner() {
	printer.Println("——————————————————————————————————————————————————————————————————")
	fmt.Print("___       ______________________________________________   __\n__ |     / /__  __ \\__  __ \\__  ____/_  ____/_  __ \\__  | / /\n__ | /| / /__  /_/ /_  /_/ /_  __/  _  /    _  / / /_   |/ / \n__ |/ |/ / _  ____/_  _, _/_  /___  / /___  / /_/ /_  /|  /  \n____/|__/  /_/     /_/ |_| /_____/  \\____/  \\____/ /_/ |_/   \n\n")
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

	printer.Done("Target:     ", database.Memory.GetString("Target"))

	if database.Memory.GetBool("Target") {
		ipTor := gohttp.TorGetIP()

		printer.Done("Proxy:      ", ipTor)
	}

	printer.Done("Started in: ", time.Now().Format(("Monday Jan 02 15:04:05 2006")))

	if names := database.Memory.GetString("Scripts List Names"); names != "" {
		var names = strings.Split(names, ",")

		for _, name := range names {
			if !scripts.Exists(name) {
				printer.Fatal("The \"" + name + "\" script does not exist")
			}
		}

		printer.Done("Loaded:     ", fmt.Sprintf("%d", len(names)), "Script(s)...")
	}

	if database.Memory.GetBool("Verbose") && database.Memory.GetBool("HTTP Options TOR") {
		printer.Danger("(Alert) Activating verbose mode together with tor mode can make the wprecon super slow.\n")
	} else if database.Memory.GetBool("Verbose") {
		printer.Danger("(Alert) Enabling verbose mode can slow down wprecon.\n")
	} else {
		printer.Println()
	}
}
