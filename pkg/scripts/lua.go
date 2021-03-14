package scripts

import (
	"path/filepath"
	"strings"

	"github.com/blackbinn/wprecon/internal/database"
	"github.com/blackbinn/wprecon/pkg/handler"
	luaNet "github.com/blackbinn/wprecon/pkg/scripts/lib/net"
	luaPrinter "github.com/blackbinn/wprecon/pkg/scripts/lib/printer"
	luaUrl "github.com/blackbinn/wprecon/pkg/scripts/lib/url"
	"github.com/blackbinn/wprecon/pkg/text"
	luaLibs "github.com/vadv/gopher-lua-libs"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

// PathScript:: this is directory scripts .lua
var (
	PathScript = "tools/scripts/"
)

type structscript struct {
	Title       string
	Author      string
	License     string
	Description string
	References  []string
	RiskLevel   string
}

type script struct {
	lstate *lua.LState
}

func NewScript() *script {
	LuaNewState := lua.NewState()

	LuaNewState.PreloadModule("url", luaUrl.Loader)
	LuaNewState.PreloadModule("printer", luaPrinter.Loader)
	LuaNewState.PreloadModule("net", luaNet.Loader)

	luaLibs.Preload(LuaNewState)

	LuaNewState.SetGlobal("tor_url", lua.LString("http://127.0.0.1:9080"))

	return &script{lstate: LuaNewState}
}

func (s *script) UseScript(n string) (sttst *structscript) {
	defer handler.HandlerErrorLuaScripts()

	if _, has := text.ContainsSliceString(AllScripts(), n); has {
		path := strings.ReplaceAll(n, ".lua", "")
		path = PathScript + path + ".lua"

		if err := s.lstate.DoFile(path); err != nil {
			panic(err)
		}

		if err := gluamapper.Map(s.lstate.GetGlobal("script").(*lua.LTable), &sttst); err != nil {
			panic(err)
		}
	} else {
		panic("The " + n + " script does not exist")
	}

	return
}

func AllScripts() []string {
	defer handler.HandlerErrorLuaScripts()

	var list []string

	scriptsfileslist, err := filepath.Glob(PathScript + "*.lua")

	if err != nil {
		panic(err)
	}

	for _, files := range scriptsfileslist {
		files = strings.ReplaceAll(files, PathScript, "")
		files = strings.ReplaceAll(files, ".lua", "")

		list = append(list, files)
	}

	return list
}

func Exists(n string) bool {
	if _, has := text.ContainsSliceString(AllScripts(), n); has {
		return true
	}

	return false
}

func GetInfos(n string) (sttst *structscript) {
	defer handler.HandlerErrorLuaScripts()

	if _, has := text.ContainsSliceString(AllScripts(), n); has {
		lstate := lua.NewState()

		path := strings.ReplaceAll(n, ".lua", "")
		path = PathScript + path + ".lua"

		lstate.PreloadModule("url", luaUrl.Loader)
		lstate.PreloadModule("printer", luaPrinter.Loader)
		lstate.PreloadModule("net", luaNet.Loader)

		luaLibs.Preload(lstate)

		lstate.SetGlobal("tor_url", lua.LString("http://127.0.0.1:9080"))

		if err := lstate.DoFile(path); err != nil {
			panic(err)
		}

		if err := gluamapper.Map(lstate.GetGlobal("script").(*lua.LTable), &sttst); err != nil {
			panic(err)
		}
	} else {
		panic("The " + n + " script does not exist")
	}

	return
}

func (s *script) Run() {
	defer s.lstate.Close()
	defer handler.HandlerErrorLuaScripts()

	err := s.lstate.CallByParam(lua.P{
		Fn:      s.lstate.GetGlobal("main"),
		NRet:    0,
		Protect: true,
	}, lua.LString(database.Memory.GetString("Target")))

	if err != nil {
		panic(err)
	}
}
