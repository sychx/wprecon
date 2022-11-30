package verbose

import (
	database "github.com/AngraTeam/wprecon/internal/memory"
	"github.com/AngraTeam/wprecon/internal/printer"
)

func Println(t string) {
    if database.Memory.GetBool("Verbose") {
        printer.Verbose(t)
    }
}