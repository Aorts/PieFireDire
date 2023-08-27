package httputil

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type HTTPPostRequestFunc func(reqBody interface{}, requestRef string) ([]byte, error)

func InitHttpClient(timeout time.Duration, maxIdleConns, maxIdleConnsPerHost, maxConnsPerHost int) *http.Client {
	certPool := x509.NewCertPool()
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            certPool,
				InsecureSkipVerify: true,
			},
			MaxIdleConns:        maxIdleConns,
			MaxIdleConnsPerHost: maxIdleConnsPerHost,
			MaxConnsPerHost:     maxConnsPerHost,
		},
	}
	return client
}

type HTTPGetFunc func() ([]byte, error)

func NewHttpGet(client *http.Client) HTTPGetFunc {
	return func() ([]byte, error) {

		req, err := http.NewRequest(http.MethodGet, "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text", nil)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to New http Request")
		}
		res, err := client.Do(req)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to New http Request")
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to New http Request ioutil")
		}
		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("%s", body)
		}
		return body, nil
	}
}
