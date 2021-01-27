package scripts

import (
	"fmt"
	"log"
	"path/filepath"

	. "github.com/blackbinn/wprecon/cli/config"
	"github.com/blackbinn/wprecon/pkg/printer"
	luaNet "github.com/blackbinn/wprecon/pkg/scripts/pkg/net"
	luaPrinter "github.com/blackbinn/wprecon/pkg/scripts/pkg/printer"
	luaUrl "github.com/blackbinn/wprecon/pkg/scripts/pkg/url"
	luaLibs "github.com/vadv/gopher-lua-libs"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

type structscript struct {
	Title       string
	Author      string
	License     string
	Description string
	References  []string
	RiskLevel   string
}

func Run(l *lua.LState) {
	err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("run"),
		NRet:    0,
		Protect: true,
	}, lua.LString(InfosWprecon.Target))

	if err != nil {
		printer.Fatal(err)
	}

	defer l.Close()
}

func Initialize(script string) (*lua.LState, *structscript) {
	var structscript structscript

	PATHScript := fmt.Sprintf("tools/scripts/%s.lua", script)

	LuaNewState := lua.NewState()

	LuaNewState.PreloadModule("url", luaUrl.Loader)
	LuaNewState.PreloadModule("printer", luaPrinter.Loader)
	LuaNewState.PreloadModule("net", luaNet.Loader)

	luaLibs.Preload(LuaNewState)

	if err := LuaNewState.DoFile(PATHScript); err != nil {
		printer.Fatal(err)
	}

	LuaNewState.SetGlobal("tor_url", lua.LString("http://127.0.0.1:9080"))

	if err := gluamapper.Map(LuaNewState.GetGlobal("script").(*lua.LTable), &structscript); err != nil {
		printer.Fatal(err)
	}

	return LuaNewState, &structscript
}

func AllScripts() []string {
	scriptsfileslist, err := filepath.Glob("tools/scripts/*.lua")

	if err != nil {
		log.Fatal(err)
	}

	return scriptsfileslist
}
