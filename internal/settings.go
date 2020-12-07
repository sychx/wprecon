package internal

import (
	"fmt"
	"time"

	"github.com/blkzy/wpgo/printer"
)

// Sbanner :: A banner that will only be executed if the scan is started correctly.
func SBanner(target string) {
	Banner()
	printer.Done("Target:", target)
	printer.Done("Starting:", time.Now().Format(("02/jan/2006 15:04:05")))
}

// Banner :: A simple banner.
func Banner() {
	fmt.Println("——————————————————————————————————————————————————")
	fmt.Printf(`
___       _______________________________ 
__ |     / /__  __ \_  ___/_  ____/_  __ \
__ | /| / /__  /_/ /____ \_  / __ _  / / /
__ |/ |/ / _  ____/____/ // /_/ / / /_/ / 
____/|__/  /_/     /____/ \____/  \____/  
`)
	fmt.Println("Github: ", "https://github.com/blackcrw/wpsgo")
	fmt.Println("Version: ", "0.0.1a")
	fmt.Println("——————————————————————————————————————————————————")
}
