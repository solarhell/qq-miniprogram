package swan_miniprogram

import "errors"

var (
	ErrNotAllowEmptyParam = errors.New("param cannot be empty")
	ErrConnectBaiduServer = errors.New("err connecet Baidu server")
)
