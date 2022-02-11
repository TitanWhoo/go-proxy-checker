package core

import "testing"

func TestCheckProxy(t *testing.T) {
	CheckProxy("socks5://127.0.0.1:7890")
}
