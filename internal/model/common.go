package model

import "github.com/mufanh/easyagent/pkg/errcode"

type BaseResponse struct {
	errcode.Error
}

func (s *BaseResponse) SetErr(err *errcode.Error) {
	s.Error = *err
}

func (s *BaseResponse) SetBizErr(err error) {
	s.Error = *errcode.NewBizErrorWithErr(err)
}

func (s *BaseResponse) Err() errcode.Error {
	return s.Error
}
