package http

import (
	"net/url"
)

/* This function will be used for URL validation */
func ThisIsURL(URL string) (error) {
    var _, err = url.ParseRequestURI(URL)

    if err != nil {
        return err
    }

    return nil
}
