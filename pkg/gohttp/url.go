package gohttp

import (
	"net"
	"net/url"

	"github.com/blackcrw/wprecon/pkg/handler"
)

// IsURL :: This function will be used for URL validation
func IsURL(URL string) bool {
	defer handler.HandlerErrorURL()

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

	//_, err = net.LookupHost(uri.Host)

	//if err != nil {
	//	return false, err
	//}

	return true
}

// GetHost ::
func GetHost(URL string) (string, error) {
	uri, err := url.ParseRequestURI(URL)

	if err != nil {
		return "", err
	}

	_, err = net.LookupHost(uri.Host)

	if err != nil {
		return "", err
	}

	return uri.Host, nil
}
