package banner

import (
	"fmt"
	"time"

	"github.com/blackcrw/wprecon/internal/config"
	"github.com/blackcrw/wprecon/internal/database"
	"github.com/blackcrw/wprecon/internal/net"
	"github.com/blackcrw/wprecon/internal/printer"
)

func Banner() {
	printer.Println("——————————————————————————————————————————————————————————————————————")
	fmt.Print(printer.BLACK+"___       ______________________________________________   __\n"+printer.BLACK+"__ |     / /__  __ \\__  __ \\__  ____/_  ____/_  __ \\__  | / /\n"+printer.BLUE+"__ | /| / /__  /_/ /_  /_/ /_  __/  _  /    _  / / /_   |/ / \n"+printer.BLUE+"__ |/ |/ / _  ____/_  _, _/_  /___  / /___  / /_/ /_  /|  /  \n"+printer.CYAN+"____/|__/  /_/     /_/ |_| /_____/  \\____/  \\____/ /_/ |_/  "+printer.RESET+config.GetConfig().App.Version+"\n\n")
	printer.Warning("Github:\t"+responsive_spaces(),"https://github.com/blackcrw/wprecon")
	printer.Println("——————————————————————————————————————————————————————————————————————")
}

func SBanner() {
	Banner()

	printer.Done("Target:\t"+responsive_spaces(), database.Memory.GetString("Options URL"))

	if database.Memory.GetBool("HTTP Options TOR") { printer.Done("Proxy:\t"+responsive_spaces(), net.TorGetIP()) }

	if database.Memory.GetBool("Options Verbose") {
		if database.Memory.GetBool("HTTP Options Random Agent") { printer.Done("Random Agent:   ON") } else { printer.Danger("Random Agent:   OFF") }
		if seconds := database.Memory.GetInt("HTTP Time Sleep"); seconds != 0 { printer.Done("Sleep Requests: "+fmt.Sprint(seconds)+"s") }
	}

	printer.Done("Started in:\t"+responsive_spaces(), time.Now().Format(("Monday Jan 02 15:04:05 2006"))+"\n")
}

func responsive_spaces() string {
	if database.Memory.GetBool("Options Verbose") {
		return "   "
	}

	return ""
}