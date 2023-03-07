package utils

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/DVKunion/collar/pkg/config"
)

func Get(url string) (*resty.Response, error) {
	client := newRequest(config.SingleConfig.Token)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("http request error: " + resp.Status())
	}
	return resp, nil
}

func newRequest(token string, ctx ...context.Context) *resty.Request {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetHeader("X-Ca-Token", token)
	req := client.R()
	if len(ctx) > 0 {
		req.SetContext(ctx[0])
	}

	return req
}
