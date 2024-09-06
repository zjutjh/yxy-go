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

// login err
const (
	ErrTokenInvalid         Code = iota + 110001 // Token无效
	ErrCaptchaInvalid                            // 图片验证码已失效
	ErrCaptchaWrong                              // 图片验证码错误
	ErrDeviceIDInconsistent                      // deviceId不一致
	ErrPhoneNumWrong                             // 手机号格式错误
	ErrSendLimit                                 // 短信发送超限
)
