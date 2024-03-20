package consts

const (
	StatusError        = 0 // 请求失败
	StatusOK           = 1 // 请求成功
	StatusUnauthorized = 2 // Token错误
	StatusLimiterError = 3 // 频繁请求错误
)
