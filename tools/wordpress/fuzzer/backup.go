package fuzzer

import (
	"time"

	"github.com/blackcrw/wprecon/pkg/gohttp"
	"github.com/blackcrw/wprecon/pkg/printer"
	"github.com/blackcrw/wprecon/pkg/wordlist"
)

// Backup ::
type Backup struct {
	HTTP    *gohttp.HTTPOptions
	Verbose bool
}

// Run ::
func (options *Backup) Run() {
	newtopline := printer.NewTopLine(":: Backup file/directory fuzzer active! ::")

	for _, path := range [...]string{"", "wp-content/", "wp-admin/"} {
		for _, directory := range wordlist.BackupFiles {
			go func(directory string) {
				options.HTTP.URL.Directory = path + directory

				response, err := gohttp.HTTPRequest(options.HTTP)

				if err != nil {
					printer.Fatal(err)
				}

				if response.StatusCode == 200 {
					newtopline.DownLine()

					printer.Done("Status Code: 200", "URL:", response.URL.Full)
				} else if response.StatusCode == 403 {
					newtopline.DownLine()

					printer.Warning("Status Code: 403", "URL:", response.URL.Full)
				}
			}(directory)

			time.Sleep(time.Millisecond * 100)
		}
	}

	if newtopline.Count <= 0 {
		newtopline.Danger(":: No backup files/directories were found. ::")
	}

	printer.Println()
}
