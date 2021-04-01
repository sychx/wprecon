package handler

import (
	"fmt"
	"strings"

	"github.com/blackbinn/wprecon/internal/pkg/printer"
)

func ErrorTorProxy() {
	if recovered := recover(); recovered != nil {
		recoveredS := fmt.Sprintf("%s", recovered)

		if strings.Contains(recoveredS, "proxyconnect tcp: dial tcp 127.0.0.1:9080: connect: connection refused") {
			printer.Fatal("Connection Refused, the tor with the command: \"tor --HTTPTunnelPort 9080\"\n")
		} else {
			printer.Danger(recoveredS).Endl()
		}
	}
}

func ErrorURL() {
	if recovered := recover(); recovered != nil {
		recoveredS := fmt.Sprintf("%s", recovered)

		printer.Fatal(recoveredS + "\n")
	}
}

func ErrorGetVuln() {
	if recovered := recover(); recovered != nil {
		recoveredS := fmt.Sprintf("%s", recovered)

		if strings.Contains(recoveredS, "dial tcp 144.217.235.104:8777: connect: connection refused") {
			printer.Danger("Connection refused to API").Endl()
		} else {
			printer.Danger(recoveredS).Endl()
		}
	}
}

func ErrorLuaScripts() {
	if recovered := recover(); recovered != nil {
		recoveredS := fmt.Sprintf("%s", recovered)

		if strings.Contains(recoveredS, "This script does not exist.") {
			printer.Fatal(recoveredS + "\n")
		} else {
			printer.Danger(recoveredS).Endl()
		}
	}
}
