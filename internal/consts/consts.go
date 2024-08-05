package consts

var (
	//for user
	CtxUserId     = "CtxUserId"
	CtxUserName   = "CtxUserName"
	CtxUserAvatar = "CtxUserAvatar"
	CtxUserSex    = "CtxUserSex"
	CtxUserSign   = "CtxUserSign"
	CtxUserStatus = "CtxUserStatus"

	//统一管理错误提示
	CodeMissingParameterMsg = "请检查是否缺少参数"
	ErrLoginFaulMsg         = "登录失败，账号或密码错误"
	ErrSecretAnswerMsg      = "密保问题不正确"
)
