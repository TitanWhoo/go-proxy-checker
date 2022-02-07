package core

import (
	"github.com/parnurzeal/gorequest"
	"net/url"
	"strings"
)

var CheckURL = "https://httpbin.org/get"

func CheckProxy(proxy string) bool {
	if strings.TrimSpace(proxy) == "" {
		return false
	}
	proxyURL := url.URL{Scheme: "http", Host: proxy}
	// get resources from pool and release after operations
	request := requestPool.Get().(*gorequest.SuperAgent)
	resp := resultPool.Get().(map[string]interface{})
	defer requestPool.Put(request)
	defer resultPool.Put(resp)
	// do the Request
	_, _, errors := request.Proxy(proxyURL.String()).Get(CheckURL).EndStruct(&resp)
	if errors != nil {
		return false
	}
	return resp["origin"] == proxyURL.Hostname()
}
