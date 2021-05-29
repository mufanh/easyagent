package model

import "github.com/mufanh/easyagent/pkg/errcode"

type baseResponse struct {
	errcode.Error
}

func (s *baseResponse) SetErr(err *errcode.Error) {
	s.Error = *err
}

func (s *baseResponse) SetBizErr(err error) {
	s.Error = *errcode.NewBizErrorWithErr(err)
}

func (s *baseResponse) Err() errcode.Error {
	return s.Error
}
