package fuzzing

import (
	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/internal/pkg/gohttp"
	"github.com/blackbinn/wprecon/internal/pkg/wordlist"
)

func BackupFile(channel chan *gohttp.Response) {
	var target = database.Memory.GetString("Target")
	var paths = [4]string{database.Memory.GetString("HTTP wp-content"), "wp-includes/", "wp-uploads/"}

	for _, directory := range paths {
		for _, file := range wordlist.BackupFiles {
			go gohttp.SimpleRequestGoroutine(channel, target, directory, file)
		}
	}
}
