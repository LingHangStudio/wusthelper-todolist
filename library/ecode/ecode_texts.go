package ecode

func InitEcodeText() {
	texts := map[Code]string{}

	// 通用ecode文案
	texts[ServerErr] = "服务器开小差了..."
	texts[NetworkErr] = "网络异常"
	texts[ParamWrong] = "参数无效"
	texts[FuncNotImpl] = "接口未实现或已弃用"
	texts[DaoOperationErr] = "dao操作错误"
	texts[TokenInvalid] = "token无效"
	texts[DataNotExists] = "数据不存在"

	Register(texts)
}
