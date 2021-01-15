package printer

import (
	"fmt"
	"io"
	"os"
	"syscall"

	color "github.com/logrusorgru/aurora" // This is color lib
	lua "github.com/yuin/gopher-lua"
)

// Required :: Constant with the word "required" in red.
var Required = color.Red("(Required)").Bold().String()
var stdout = *os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
var stderr = *os.NewFile(uintptr(syscall.Stdout), "/dev/stderr")

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)

	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"done":    done,
	"danger":  danger,
	"warning": warning,
	"fatal":   fatal,
}

// Done ::
func done(L *lua.LState) int {
	var prefix = color.Green("[✔]").String()

	_, err := io.WriteString(&stdout, prefix+" "+L.ToString(1)+"\n")

	if err != nil {
		panic(err)
	}

	return 0
}

// Danger ::
func danger(L *lua.LState) int {
	var prefix = color.Red("[✗]").String()

	_, err := io.WriteString(&stdout, prefix+" "+L.ToString(1)+"\n")

	if err != nil {
		panic(err)
	}

	return 0
}

// Warning ::
func warning(L *lua.LState) int {
	var prefix = color.Yellow("[!]").String()

	_, err := io.WriteString(&stdout, prefix+" "+L.ToString(1)+"\n")

	if err != nil {
		panic(err)
	}

	return 0
}

// Fatal ::
func fatal(L *lua.LState) int {
	var prefix = color.Red("[!]").String()

	fmt.Fprint(&stderr, prefix, " ")
	fmt.Fprintln(&stderr, L.ToString(1))

	os.Exit(0)

	return 0
}
