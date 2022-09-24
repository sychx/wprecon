package http

/*
A simple request function;
the requests made by this function will always be in the GET method,
and by default it comes with some predefined parameters, such as random agent and the tls certificate disabled.
*/
func Request(URL string) (*EntityResponse, error) {
	var http = NewHTTP(URL)

	http.SetMethod("GET")
	http.OnRandomUserAgent(true)
	http.OnTLSCertificateVerify(false)

	var response, err = Do(http)
	if err != nil { return nil, err }

	return response, nil
}
