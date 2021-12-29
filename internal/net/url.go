package net

import (
	"net"
	"net/url"

	"github.com/blackcrw/wprecon/internal/printer"
)

/* This function will be used for URL validation */
func ThisIsURL(URL string) bool {
	uri, err := url.ParseRequestURI(URL)

	if err != nil {
		printer.Fatal(err)
	}

	switch uri.Scheme {
	case "http":
	case "https":
	default:
		printer.Fatal("Invalid scheme")
	}

	return true
}

func ThisIsHostValid(URL string) bool {
	var uri, err = url.ParseRequestURI(URL)

	if err != nil {
		printer.Fatal(err)
	}

	_, err = net.LookupHost(uri.Host)

	if err != nil {
		printer.Fatal(err)
	}

	return true
}