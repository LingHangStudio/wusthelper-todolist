package ecode

// 抄的旧版状态码
// All common ecode
var (
	OK = New(10000) // 正确

	ServerErr     = New(11111)  // 服务器错误
	NetworkErr    = New(10_002) // 网络错误
	FuncNotImpl   = New(10_004) // api未实现
	ParamWrong    = New(10001)  // 参数错误
	TokenInvalid  = New(21002)  // token无效
	DataNotExists = New(10_100) // token无效

	DaoOperationErr = New(10_101)
)
