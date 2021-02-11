package main

import (
	"runtime"

	"github.com/blackbinn/wprecon/cli"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cli.Execute()
}
