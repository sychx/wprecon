package http

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

func NewHTTP(URL string) (httpOptions *_EntityOptionsHttp) {
    httpOptions = &_EntityOptionsHttp{
        url:         URL,
        method:      "GET",
        userAgent:   "WPrecon - Wordpress Recon (Vulnerability Recon)",
        data:        nil,
        contentType: "text/html; charset=UTF-8",
    }

    return
}

func (model *_EntityOptionsHttp) OnRandomUserAgent(status bool) *_EntityOptionsHttp {
    if status { model.userAgent = randomUserAgent() }

    return model
}

func (model *_EntityOptionsHttp) OnTLSCertificateVerify(status bool) *_EntityOptionsHttp { model.tlsCertificateVerify = status; return model }
func (model *_EntityOptionsHttp) SetUserAgent(agent string) *_EntityOptionsHttp { model.userAgent = agent; return model }
func (model *_EntityOptionsHttp) SetMethod(method string) *_EntityOptionsHttp { model.method = method; return model }
func (model *_EntityOptionsHttp) SetSleep(tm int) *_EntityOptionsHttp { model.timeSleep = time.Duration(tm); return model }

func Do(model *_EntityOptionsHttp) (*EntityResponse, error) {
    var (
        err error
        bodyRaw []byte
        request *http.Request
        response *http.Response
        
        client = &http.Client{
            // CheckRedirect: model.redirect,
            Transport: &http.Transport{
                // Proxy:             model.proxy,
                DisableKeepAlives: true,
                TLSClientConfig: &tls.Config{
                    InsecureSkipVerify: model.tlsCertificateVerify,
                },
            },
        }
    )

    request, err = http.NewRequest(model.method, model.url, nil)
    if err != nil { return nil, err }

    request.Header.Set("User-Agent", model.userAgent)
    request.Header.Set("Content-Type", model.contentType)

    response, err = client.Do(request)
    if err != nil { return nil, err }

    bodyRaw, err = io.ReadAll(response.Body)
    if err != nil { return nil, err }

    var entity = &EntityResponse{
        Raw:      string(bodyRaw),
        URL:      response.Request.URL,
        Response: response,
    }

    if sleep := model.timeSleep; sleep != 0 {
        time.Sleep(sleep * time.Second)
    }

    return entity, nil
}