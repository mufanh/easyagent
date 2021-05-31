package errcode

import (
	"fmt"
)

type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg,omitempty"`
	Details []string `json:"details,omitempty"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{Code: code, Msg: msg}
}

func NewBizErrorWithErr(err error) *Error {
	return &Error{Code: BizError.Code, Msg: err.Error()}
}

func NewBizErrorWithMsg(msg string) *Error {
	return &Error{Code: BizError.Code, Msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code, e.Msg)
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	e.Details = []string{}
	for _, d := range details {
		e.Details = append(e.Details, d)
	}
	return e
}
