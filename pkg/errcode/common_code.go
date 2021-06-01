package errcode

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(10000001, "服务内部错误")
	InvalidParams = NewError(10000002, "入参错误")
	BizError      = NewError(10000003, "业务错误")
)
