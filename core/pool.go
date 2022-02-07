package core

import (
	"crypto/tls"
	"github.com/parnurzeal/gorequest"
	"sync"
	"time"
)

var agent *gorequest.SuperAgent
var requestPool = sync.Pool{
	New: func() interface{} {
		return agent.Clone()
	},
}
var resultPool = sync.Pool{
	New: func() interface{} {
		return map[string]interface{}{}
	},
}

func init() {
	agent = gorequest.New()
	agent.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	agent.Timeout(15 * time.Second)
}
