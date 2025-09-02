package yxyClient

import (
	urllib "net/url"
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
	parsedURL, err := urllib.Parse(url)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrHttpClient, "invalid url")
	}
	// 根据域名判断是否需要添加 sign
	if parsedURL.Hostname() == "compus.xiaofubao.com" {
		sign := Sign(req)
		headers["sign"] = sign
	}
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
