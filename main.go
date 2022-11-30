/*
Está aplicação implementa a mentodologia SOLID.
*/

package main

import (
	"github.com/AngraTeam/wprecon/cli"
	_ "github.com/AngraTeam/wprecon/cli"
)

func main() {
<<<<<<< Updated upstream
	var URL = ""

	interesting.WordpressCheck(URL, "")
	interesting.DirectoryPlugins(URL)
	interesting.DirectoryThemes(URL)
	interesting.DirectoryUploads(URL)

//	var request, _ = http.Request(URL)

//	fmt.Println(plugins.IndexSourceCodeBody("/wp-content", request.Raw))
//	fmt.Println(plugins.IndexSourceCodeBodyVersion("/wp-content", request.Raw))
}
=======
    cli.Start()
}
>>>>>>> Stashed changes
