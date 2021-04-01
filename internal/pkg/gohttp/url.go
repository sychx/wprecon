package gohttp

import (
	"net"
	"net/url"

	"github.com/blackbinn/wprecon/internal/pkg/handler"
)

// IsURL :: This function will be used for URL validation
func IsURL(URL string) bool {
	defer handler.ErrorURL()

	uri, err := url.ParseRequestURI(URL)

	if err != nil {
		panic(err)
	}

	switch uri.Scheme {
	case "http":
	case "https":
	default:
		panic("Invalid scheme")
	}

	/*
		if _, err = net.LookupHost(uri.Host); err != nil {
			panic(err)
		}
	*/

	return true
}

// GetHost ::
func GetHost(URL string) (string, error) {
	uri, err := url.ParseRequestURI(URL)

	if err != nil {
		return "", err
	}

	if _, err = net.LookupHost(uri.Host); err != nil {
		return "", err
	}

	return uri.Host, nil
}
