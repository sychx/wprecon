package fuzzing

import (
	"time"

	. "github.com/blackcrw/wprecon/cli/config"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/pkg/wordlist"
	"github.com/blackcrw/wprecon/tools/wordpress/extensions"
)

func BackupFile() {
	printer.Warning(":: Backup file/directory fuzzer active! ::")

	done := false

	for _, directory := range [...]string{"", "wp-content/", "wp-admin/", "wp-includes/"} {
		for _, file := range wordlist.BackupFiles {
			go func(file string) {
				response := extensions.SimpleRequest(InfosWprecon.Target, directory+file)

				if response.StatusCode == 200 {
					printer.Done("Status Code: 200", "URL:", response.URL.Full)
					done = true
				} else if response.StatusCode == 403 {
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
