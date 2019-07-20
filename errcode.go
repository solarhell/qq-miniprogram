package qq_miniprogram

import "errors"

var (
	ErrNotAllowEmptyParam   = errors.New("param cannot be empty")
	ErrConnectTencentServer = errors.New("err connect Tencent server")
)
