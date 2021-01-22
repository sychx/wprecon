package scripts

import (
	"fmt"
	"log"
	"path/filepath"

	. "github.com/blackcrw/wprecon/cli/config"
	"github.com/blackcrw/wprecon/pkg/printer"
	luaNet "github.com/blackcrw/wprecon/pkg/scripts/pkg/net"
	luaPrinter "github.com/blackcrw/wprecon/pkg/scripts/pkg/printer"
	luaUrl "github.com/blackcrw/wprecon/pkg/scripts/pkg/url"
	"github.com/blackcrw/wprecon/pkg/text"
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

	if _, has := text.ContainsSliceString(AllScripts(), script); has == true {
		PATHScript := fmt.Sprintf("scripts/%s.lua", script)

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

	return &lua.LState{}, &structscript
}

func AllScripts() []string {
	scriptsfileslist, err := filepath.Glob("scripts/*.lua")

	if err != nil {
		log.Fatal(err)
	}

	return scriptsfileslist
}
