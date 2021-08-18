package utils

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	HTTPS = "https"
	HTTP  = "http"
)

func DoHttp(req *http.Request, w http.ResponseWriter) {
	var (
		response *http.Response
		err      error
	)
	if strings.HasPrefix(strings.ToLower(req.URL.Scheme[:3]), HTTPS) {
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
		w.Write([]byte(err.Error()))
		return
	}
	for k, v := range response.Header {
		w.Header().Set(k, v[0])
	}
	bytes, _ := ioutil.ReadAll(req.Body)
	w.Write(bytes)

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
