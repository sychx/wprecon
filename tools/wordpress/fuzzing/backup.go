package fuzzing

import (
	"time"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/pkg/gohttp"
	"github.com/blackbinn/wprecon/pkg/printer"
	"github.com/blackbinn/wprecon/pkg/wordlist"
)

func BackupFile() {
	printer.Warning(":: Backup file/directory fuzzer active! ::")

	done := false

	for _, directory := range [...]string{"", InfosWprecon.WPContent, "wp-includes/", "wp-uploads/"} {
		for _, file := range wordlist.BackupFiles {
			go func(file string) {
				response := gohttp.SimpleRequest(InfosWprecon.Target, directory+file)

				if response.Response.StatusCode == 200 {
					printer.Done("Status Code: 200", "URL:", response.URL.Full)
					done = true
				} else if response.Response.StatusCode == 403 {
					printer.Warning("Status Code: 403", "URL:", response.URL.Full)
					done = true
				}
			}(file)

			time.Sleep(time.Millisecond * 100)
		}
	}

	if !done {
		printer.Danger(":: No backup files/directories were found. ::")
	}
}
