package core

import (
	"github.com/parnurzeal/gorequest"
	"strings"
)

var CheckURL = "https://httpbin.org/get"

func CheckProxy(proxy string) bool {
	if strings.TrimSpace(proxy) == "" {
		return false
	}
	// no protocol, add //
	if !strings.Contains(proxy, "//") {
		proxy = "//" + proxy
	}
	// get resources from pool and release after operations
	request := requestPool.Get().(*gorequest.SuperAgent)
	resp := resultPool.Get().(map[string]interface{})
	defer requestPool.Put(request)
	defer resultPool.Put(resp)
	// do the Request
	_, _, errors := request.Proxy(proxy).Get(CheckURL).EndStruct(&resp)
	if errors != nil {
		return false
	}
	return strings.Contains(proxy, resp["origin"].(string))
}
