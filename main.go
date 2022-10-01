/*
Está aplicação implementa a mentodologia SOLID.
*/

package main

import (
	"fmt"

	"github.com/AngraTeam/wprecon/internal/http"
	"github.com/AngraTeam/wprecon/tools/finders/plugins"
	"github.com/AngraTeam/wprecon/tools/interesting"
)

func main() {
	var URL = "https://www.poli.usp.br/"

	interesting.WordpressCheck(URL, "")
	interesting.DirectoryPlugins(URL)
	interesting.DirectoryThemes(URL)
	interesting.DirectoryUploads(URL)

	var request, _ = http.Request(URL)

	fmt.Println(plugins.IndexSourceCodeBody("/wp-content", request.Raw))
	fmt.Println(plugins.IndexSourceCodeBodyVersion("/wp-content", request.Raw))
}
