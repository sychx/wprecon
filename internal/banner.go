package internal

import (
	"fmt"
	"time"

	"github.com/blkzy/wpsgo/pkg/printer"
)

// Sbanner :: A banner that will only be executed if the scan is started correctly.
func SBanner(target string) {
	Banner()
	printer.Done("Target:", target)
	printer.Done("Starting:", time.Now().Format(("02/jan/2006 15:04:05")))
}

// Banner :: A simple banner.
func Banner() {
	fmt.Print("——————————————————————————————————————————————————")
	fmt.Printf(`
___       _______________________________ 
__ |     / /__  __ \_  ___/_  ____/_  __ \
__ | /| / /__  /_/ /____ \_  / __ _  / / /
__ |/ |/ / _  ____/____/ // /_/ / / /_/ / 
____/|__/  /_/     /____/ \____/  \____/  
`)
	fmt.Print("Github: ", "https://github.com/blackcrw/wpsgo\n")
	fmt.Print("Version: ", "0.0.1a\n")
	fmt.Print("——————————————————————————————————————————————————\n")
}
