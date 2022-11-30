package plugins

import (
	"regexp"

	db "github.com/AngraTeam/wprecon/internal/memory"
	"github.com/AngraTeam/wprecon/internal/printer"
)

func PassiveVersionSearch() {
    var raw = db.Memory.GetString("HTTP Index Raw")
    var pathcontent = db.Memory.GetString("HTTP wp-content")

    var regex = regexp.MustCompile(pathcontent+`/plugins/(.*?)/.*?[.css|.js]`)
    for _, submatch := range regex.FindAllStringSubmatch(raw, -1) {

        printer.Danger(submatch)
    }
}