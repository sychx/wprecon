package gohttp

import (
	"errors"
	"net"
	"net/url"
)

/*
	I was seriously in doubt as to how this validation would do. So I went after examples of validators... until I found this code... and I found it perfect!
	:: https://stackoverflow.com/questions/51069484/url-validation-seems-broken
*/

// URL ::
type URL interface {
	IsURL() (bool, error)
}

// IsURL :: This function will be used for URL validation
func IsURL(URL string) (bool, error) {
	// Check it's an Absolute URL or absolute path
	uri, err := url.ParseRequestURI(URL)
	if err != nil {
		return false, err
	}

	// Check it's an acceptable scheme
	switch uri.Scheme {
	case "http":
	case "https":
	default:
		return false, errors.New("Invalid scheme")
	}

	// Check it's a valid domain name
	_, err = net.LookupHost(uri.Host)
	if err != nil {
		return false, err
	}

	return true, nil
}
