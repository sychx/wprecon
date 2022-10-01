/*
Está aplicação implementa a mentodologia SOLID.
*/

package main

import (
	"github.com/AngraTeam/wprecon/tools/interesting"
)

func main() {
	var URL = "https://www.poli.usp.br/"

	interesting.WordpressCheck(URL, "")
	interesting.DirectoryPlugins(URL)
	interesting.DirectoryThemes(URL)
	interesting.DirectoryUploads(URL)

//	var request, _ = http.Request(URL)

//	fmt.Println(plugins.IndexSourceCodeBody("/wp-content", request.Raw))
//	fmt.Println(plugins.IndexSourceCodeBodyVersion("/wp-content", request.Raw))
}
