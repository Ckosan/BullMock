package utils

import (
	"crypto/tls"
	"net/http"
)

const (
	HTTPS = "https://"
	HTTP  = "http://"
)

func DoHttp(req *http.Request) (error, *http.Response) {
	var (
		response *http.Response
		err      error
	)

	if req.TLS != nil {
		response, err = httpsReq(req)
	} else {
		response, err = httpReq(req)
	}

	defer func() {
		if response != nil {
			response.Body.Close()
		}
	}()

	if err != nil {
		return err, nil
	}
	return nil, response
}

func httpReq(req *http.Request) (*http.Response, error) {
	client := http.DefaultClient
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func httpsReq(r *http.Request) (*http.Response, error) {
	tls11Transport := &http.Transport{
		MaxIdleConnsPerHost: 10,
		TLSClientConfig: &tls.Config{
			MaxVersion:         tls.VersionTLS13,
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tls11Transport,
	}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
