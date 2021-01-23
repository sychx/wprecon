package handler

import (
	"fmt"
	"strings"

	"github.com/blackcrw/wprecon/pkg/printer"
)

// HandlerErrorTorProxy ::
func HandlerErrorTorProxy() {
	if recovered := recover(); recovered != nil {
		recoveredString := fmt.Sprintf("%s", recovered)

		if strings.Contains(recoveredString, "proxyconnect tcp: dial tcp 127.0.0.1:9080: connect: connection refused") {
			printer.Fatal("Connection Refused, the tor with the command: \"tor --HTTPTunnelPort 9080\"")
		} else {
			printer.Danger(recoveredString)
		}
	}
}

// HandlerErrorURL ::
func HandlerErrorURL() {
	if recovered := recover(); recovered != nil {
		recoveredString := fmt.Sprintf("%s", recovered)

		printer.Fatal(recoveredString)
	}
}

func HandlerErrorGetVuln() {
	if recovered := recover(); recovered != nil {
		recoveredString := fmt.Sprintf("%s", recovered)

		if strings.Contains(recoveredString, "dial tcp 144.217.235.104:8777: connect: connection refused") {
			printer.Danger("Connection refused to API")
		} else {
			printer.Danger(recoveredString)
		}
	}
}
