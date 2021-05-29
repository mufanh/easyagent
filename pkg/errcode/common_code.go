package errcode

var (
	Success       = NewError(0, "成功")
	ServerError   = NewError(10000000, "服务内部错误")
	InvalidParams = NewError(10000001, "入参错误")
	NotFound      = NewError(10000002, "找不到")
	BizError      = NewError(10000004, "业务错误")

	TooManyRequests = NewError(10000007, "请求过多")
)
