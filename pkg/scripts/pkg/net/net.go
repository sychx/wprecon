package net

import (
	"fmt"
	"net"

	"github.com/blackbinn/wprecon/pkg/printer"
	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)

	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"lookup_ip": lookupip,
}

func lookupip(L *lua.LState) int {
	ips, err := net.LookupIP(L.ToString(1))

	if err != nil {
		printer.Fatal(err)
	}

	ip := fmt.Sprintf("%s", ips[0])

	L.Push(lua.LString(ip))

	return 1
}
