package gohttp

import (
	"crypto/tls"
	"net/http"

	"github.com/blkzy/wpsgo/pkg/gohttp"
)

// HttpRequest :: This function will be used for any request that is made.
func HttpRequest(httpStructs gohttp.Http) (gohttp.Result, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	request, err := http.NewRequest("GET", httpStructs.URL, nil)

	if err != nil {
		return gohttp.Result{}, err
	}

	response, err := client.Do(request)

	if err != nil {
		return gohttp.Result{}, err
	}

	request.Header.Set("User-Agent", "Wordpress Security Go (GoHttp 0.1.0)")
	if httpStructs.Options.RandomUserAgent == true {
		userAgent := RandomUserAgent()
		request.Header.Set("User-Agent", userAgent)
	}

	httpResult := gohttp.Result{
		URL:        request.URL.Scheme + "://" + request.URL.Host + request.URL.Path,
		StatusCode: response.StatusCode,
		UserAgent:  request.UserAgent(),
		Body:       response.Body,
	}

	return httpResult, nil
}
