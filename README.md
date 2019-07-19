# swan-miniprogram
[![Build Status](https://travis-ci.org/solarhell/mina.svg?branch=master)](https://travis-ci.org/solarhell/mina)
百度智能小程序 golang sdk


## done
登录
AccessToken(需持久化 防止超过请求限制)

## todo


## usage

### 登录
```go
package main

import (
	swan "github.com/solarhell/swan-miniprogram"
	"net/http"
	"time"
)

func main() {
	c := swan.NewClient(&http.Client{
		Timeout: 30 * time.Second,
		Transport: &swan.DebugRequestTransport{
			RequestHeader:  true,
			RequestBody:    true,
			ResponseHeader: true,
			ResponseBody:   true,
			Transport: &http.Transport{
				IdleConnTimeout: 30 * time.Second,
	        },
		},
	})

	ui, err := c.Login("appKey", "appSecret", "code")
	...
}
```
