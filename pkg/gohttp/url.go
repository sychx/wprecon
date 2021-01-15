package gohttp

import (
	"errors"
	"net"
	"net/url"
)

// IsURL :: This function will be used for URL validation
func IsURL(URL string) (bool, error) {
	uri, err := url.ParseRequestURI(URL)

	if err != nil {
		return false, err
	}

	switch uri.Scheme {
	case "http":
	case "https":
	default:
		return false, errors.New("Invalid scheme")
	}

	_, err = net.LookupHost(uri.Host)

	if err != nil {
		return false, err
	}

	return true, nil
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
