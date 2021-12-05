package banner

import (
	"fmt"
	"time"

	"github.com/blackcrw/wprecon/internal/config"
	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/printer"
	"github.com/blackcrw/wprecon/internal/update"
)

func Banner() {
	printer.Println("——————————————————————————————————————————————————————————————————")
	fmt.Print(printer.Green+"___       ______________________________________________   __\n"+printer.Black+"__ |     / /__  __ \\__  __ \\__  ____/_  ____/_  __ \\__  | / /\n"+printer.Green+"__ | /| / /__  /_/ /_  /_/ /_  __/  _  /    _  / / /_   |/ / \n"+printer.Black+"__ |/ |/ / _  ____/_  _, _/_  /___  / /___  / /_/ /_  /|  /  \n"+printer.Green+"____/|__/  /_/     /_/ |_| /_____/  \\____/  \\____/ /_/ |_/"+printer.Reset+"   \n\n")
	printer.Println("Github  :", "https://github.com/blackcrw/wprecon")

	if version_api := update.GetVersion(); config.GetConfig().App.Version != version_api {
		printer.Println("Version :", config.GetConfig().App.Version, "(Latest version: "+version_api+")")
	} else {
		printer.Println("Version :", config.GetConfig().App.Version)
	}

	printer.Println("——————————————————————————————————————————————————————————————————")
}

func SBanner() {
	Banner()

	printer.Done("Target:\t", database.Memory.GetString("Target"))

	if database.Memory.GetBool("HTTP Options TOR") {
		var ip_tor = net.TorGetIP()

		printer.Done("Proxy:\t", ip_tor)
	}

	printer.Done("Started in:\t", time.Now().Format(("Monday Jan 02 15:04:05 2006"))).Endl()

	if database.Memory.GetBool("Verbose") && database.Memory.GetBool("HTTP Options TOR") {
		printer.Warning("Activating verbose mode together with tor mode can make the wprecon super slow.\n")
	} else if database.Memory.GetBool("Verbose") {
		printer.Warning("Enabling verbose mode can slow down wprecon.\n")
	} else {
		printer.Println()
	}
}