package interesting

import (
	"strings"

	"github.com/blackcrw/wprecon/internal/http"
	. "github.com/blackcrw/wprecon/internal/memory"
)

func DirectoryPlugins(URL string) (Interesting, error) {
	var request = http.NewHTTP(URL + "/plugins/")

	request.OnRandomUserAgent(true)
	request.OnTLSCertificateVerify(false)

	var response, err = http.Do(request)

	if err != nil { return Interesting{}, err }

	if strings.Contains(response.Raw, "Index of") {
		Memory.AddInSlice("Index Of's", response.URL.String())
	}

	var entity = Interesting{
		Url:        response.URL.String(),
		Raw:        response.Raw,
		Status:     response.Response.StatusCode,
		FoundBy:    "Direct Access",
		Confidence: 0,
	}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		entity.Confidence = 100
	}

	return entity, nil
}

func DirectoryThemes(URL string) (Interesting, error) {
	var request = http.NewHTTP(URL + "/themes/")

	request.OnRandomUserAgent(true)
	request.OnTLSCertificateVerify(false)

	var response, err = http.Do(request)

	if err != nil { return Interesting{}, err }

	if strings.Contains(response.Raw, "Index of") {
		Memory.AddInSlice("Index Of's", response.URL.String())
	}

	var entity = Interesting{
		Url:        response.URL.String(),
		Raw:        response.Raw,
		Status:     response.Response.StatusCode,
		FoundBy:    "Direct Access",
		Confidence: 0,
	}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		entity.Confidence = 100
	}

	return entity, nil
}

func DirectoryUploads(URL string) (Interesting, error) {
	var request = http.NewHTTP(URL + "/uploads/")

	request.OnRandomUserAgent(true)
	request.OnTLSCertificateVerify(false)

	var response, err = http.Do(request)

	if err != nil { return Interesting{}, err }

	if strings.Contains(response.Raw, "Index of") {
		Memory.AddInSlice("Index Of's", response.URL.String())
	}

	var entity = Interesting{
		Url:        response.URL.String(),
		Raw:        response.Raw,
		Status:     response.Response.StatusCode,
		FoundBy:    "Direct Access",
		Confidence: 0,
	}

	if response.Response.StatusCode == 200 || response.Response.StatusCode == 403 {
		entity.Confidence = 100
	}

	return entity, nil
}