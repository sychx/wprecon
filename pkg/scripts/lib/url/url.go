package url

import (
	"fmt"
	"net"
	"net/url"

	"github.com/blackbinn/wprecon/pkg/printer"
	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {

	L.Push(L.SetFuncs(L.NewTable(), exports))

	return 1
}

var exports = map[string]lua.LGFunction{
	"host": gethost,
}

func gethost(L *lua.LState) int {
	uri, err := url.ParseRequestURI(L.ToString(1))

	if err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	if _, err := net.LookupHost(uri.Host); err != nil {
		printer.Danger(fmt.Sprintf("%s", err))
	}

	L.Push(lua.LString(uri.Host))

	return 1
}
