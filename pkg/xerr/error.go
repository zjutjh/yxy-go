package xerr

import (
	"fmt"
)

type ErrCode struct {
	code Code
	msg  string
}

func (e *ErrCode) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code(), e.Msg())
}

func (e *ErrCode) Code() Code {
	return e.code
}

func (e *ErrCode) Msg() string {
	return e.msg
}

func WithCode(code Code, msg string) error {
	return &ErrCode{
		code: code,
		msg:  msg,
	}
}
