# qq-miniprogram
[![Build Status](https://travis-ci.org/solarhell/mina.svg?branch=master)](https://travis-ci.org/solarhell/mina)
QQ小程序 golang sdk


## done
登录
AccessToken(需持久化 防止超过请求限制)
发送模板消息

## todo


## usage

### 登录
```go
package main

import (
	QM "github.com/solarhell/qq-miniprogram"
)

func main() {
	ui, err := QM.Login("appid", "secret", "code")
	...
}
```
