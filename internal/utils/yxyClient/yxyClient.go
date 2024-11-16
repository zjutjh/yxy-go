package yxyClient

import (
	"sync"
	"yxy-go/pkg/xerr"

	"github.com/go-resty/resty/v2"
)

var (
	client *resty.Client
	once   sync.Once
)

func initClient() {
	client = resty.New().
		SetRedirectPolicy(resty.NoRedirectPolicy()).
		SetCookieJar(nil)
}

func GetClient() *resty.Client {
	once.Do(initClient)
	return client
}

func HttpSendPost(url string, req map[string]interface{}, headers map[string]string, resp interface{}) (*resty.Response, error) {
	client := GetClient()

	r, err := client.R().
		SetHeaders(headers).
		SetBody(req).
		SetResult(&resp).
		Post(url)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	return r, nil
}
