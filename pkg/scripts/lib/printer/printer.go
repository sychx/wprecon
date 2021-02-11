package printer

import (
	"fmt"
	"io"
	"os"
	"runtime"

	color "github.com/logrusorgru/aurora"
	lua "github.com/yuin/gopher-lua"
)

var (
	stdin  = *os.Stdin
	stdout = *os.Stdout
	stderr = *os.Stderr
)

var (
	prefixDanger  = color.Red("    —").String()
	prefixDone    = color.Green("    —").String()
	prefixDefault = color.White("    —").String()
	prefixWarning = color.Yellow("    —").String()
)

func init() {
	if runtime.GOOS == "windows" {
		prefixDanger = "    —"
		prefixDone = "    —"
		prefixDefault = "    —"
		prefixWarning = "    —"
	}
}

func Loader(L *lua.LState) int {

	L.Push(L.SetFuncs(L.NewTable(), exports))

	return 1
}

var exports = map[string]lua.LGFunction{
	"done":    done,
	"danger":  danger,
	"warning": warning,
	"fatal":   fatal,
}

func done(L *lua.LState) int {
	io.WriteString(&stdout, prefixDone+" "+L.ToString(1)+"\n")

	return 0
}

func danger(L *lua.LState) int {
	io.WriteString(&stdout, prefixDanger+" "+L.ToString(1)+"\n")

	return 0
}

func warning(L *lua.LState) int {
	io.WriteString(&stdout, prefixWarning+" "+L.ToString(1)+"\n")

	return 0
}

func fatal(L *lua.LState) int {
	fmt.Fprint(&stderr, prefixDanger, " ")
	fmt.Fprintln(&stderr, L.ToString(1))

	os.Exit(0)

	return 0
}
