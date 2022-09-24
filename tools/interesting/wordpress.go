package interesting

import (
	"regexp"
	"strings"
	"sync"

	. "github.com/blackcrw/wprecon/internal/memory"
	"github.com/blackcrw/wprecon/internal/printer"
)

func WordPressVersion(raw string) string {
	var regxp = regexp.MustCompile("<meta name=\"generator\" content=\"WordPress ([0-9.-]*).*?")

	for _, sliceBytes := range regxp.FindAllSubmatch([]byte(raw), -1) {
		Memory.SetString("WordPress Version", string(sliceBytes[1][:]))
	}

	return Memory.GetString("WordPress Version")
}

func WordpressCheck(URL string, raw string) int8 {
	var (
		confidence int8
		wait_group sync.WaitGroup
		
		payloads = [4]string{
			"<meta name=\"generator content=\"WordPress",
			"<a href=\"http://www.wordpress.com\">Powered by WordPress</a>",
			"<link rel=\"https://api.wordpress.org/",
			"<link rel=\"https://api.w.org/\"",
		}
	)

	wait_group.Add(4)

	go func(){ if check, err := AdminPage(URL);        check.Confidence == 100 { confidence++ } else if err != nil { printer.Danger(err) }; wait_group.Done() }()
	go func(){ if check, err := DirectoryThemes(URL);  check.Confidence == 100 { confidence++ } else if err != nil { printer.Danger(err) }; wait_group.Done() }()
	go func(){ if check, err := DirectoryPlugins(URL); check.Confidence == 100 { confidence++ } else if err != nil { printer.Danger(err) }; wait_group.Done() }()
	go func(){ if check, err := DirectoryUploads(URL); check.Confidence == 100 { confidence++ } else if err != nil { printer.Danger(err) }; wait_group.Done() }()

	for _, payload := range payloads {
		if strings.Contains(raw, payload) {
			confidence++
		}
	}

	wait_group.Wait()

	return confidence / 8 * 100
}