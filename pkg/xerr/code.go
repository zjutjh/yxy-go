package xerr

//go:generate stringer -type Code -linecomment

type Code int

// success
const ErrSuccess Code = 0 // Success

// common err
const (
	ErrUnknown    Code = iota + 100001 // 服务异常
	ErrParam                           // 参数错误
	ErrHttpClient                      // HTTP客户端请求错误
)
