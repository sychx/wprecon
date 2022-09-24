/*
Está aplicação implementa a mentodologia SOLID.
*/

package main

import (
	"fmt"

	"github.com/blackcrw/wprecon/internal/http"
	"github.com/blackcrw/wprecon/tools/finders/plugins"
	"github.com/blackcrw/wprecon/tools/interesting"
)

func main() {
	var URL string = "https://www.poli.usp.br/"

	interesting.WordpressCheck(URL, "")
	interesting.DirectoryPlugins(URL)
	interesting.DirectoryThemes(URL)
	interesting.DirectoryUploads(URL)

	var request, _ = http.Request(URL)

	fmt.Println(plugins.IndexSourceCodeBody("/wp-content", request.Raw))
	fmt.Println(plugins.IndexSourceCodeBodyVersion("/wp-content", request.Raw))
}