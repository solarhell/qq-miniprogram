package qq_miniprogram

import "errors"

var (
	ErrNotAllowEmptyParam   = errors.New("param cannot be empty")
	ErrSignature            = errors.New("signature error")
	ErrPaddingSize          = errors.New("padding size error")
	ErrAppidMismatch        = errors.New("app id mismtch error")
	ErrDataExpire           = errors.New("data expire error")
	ErrConnectTencentServer = errors.New("err connecet Tencent server")
)
