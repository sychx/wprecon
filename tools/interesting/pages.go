package interesting

import "github.com/blackcrw/wprecon/internal/http"

func AdminPage(URL string) (Interesting, error) {
	var request = http.NewHTTP(URL + "/wp-admin/")

	request.OnRandomUserAgent(true)
	request.OnTLSCertificateVerify(false)

	var response, err = http.Do(request)

	if err != nil { return Interesting{}, err }

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

func RobotsPage(URL string) (Interesting, error) {
	var request = http.NewHTTP(URL + "/robots.txt")

	request.OnRandomUserAgent(true)
	request.OnTLSCertificateVerify(false)

	var response, err = http.Do(request)

	if err != nil { return Interesting{}, err }

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

func SiteMapPage(URL string) (Interesting, error) {
	var request = http.NewHTTP(URL + "/sitemap.xml")

	request.OnRandomUserAgent(true)
	request.OnTLSCertificateVerify(false)

	var response, err = http.Do(request)

	if err != nil { return Interesting{}, err }

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

func ReadmePage(URL string) (Interesting, error) {
	var request = http.NewHTTP(URL + "/readme.html")

	request.OnRandomUserAgent(true)
	request.OnTLSCertificateVerify(false)

	var response, err = http.Do(request)

	if err != nil { return Interesting{}, err }

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