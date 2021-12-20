package banner

import (
	"fmt"
	"time"

	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/printer"
)

func Banner() {
	printer.Println("——————————————————————————————————————————————————————————————————")
	fmt.Print(printer.Black+"___       ______________________________________________   __\n"+printer.Black+"__ |     / /__  __ \\__  __ \\__  ____/_  ____/_  __ \\__  | / /\n"+printer.Blue+"__ | /| / /__  /_/ /_  /_/ /_  __/  _  /    _  / / /_   |/ / \n"+printer.Blue+"__ |/ |/ / _  ____/_  _, _/_  /___  / /___  / /_/ /_  /|  /  \n"+printer.Cyan+"____/|__/  /_/     /_/ |_| /_____/  \\____/  \\____/ /_/ |_/"+printer.Reset+"   \n\n")
	printer.Warning("Github:\t","https://github.com/blackcrw/wprecon")
	
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
}