package consts

var (

	//for admin
	CtxAdminId      = "CtxAdminId"
	CtxAdminName    = "CtxAdminName"
	CtxAdminIsAdmin = "CtxAdminIsAdmin"
	CtxAdminRoleIds = "CtxAdminRoleIds"

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
	ResourcePermissionFail  = "没有权限操作"

	//for 登录相关
	GTokenAdminPrefix         = "Admin:" //gtoken登录 管理后台 前缀区分
	GTokenFrontendPrefix      = "User:"  //gtoken登录 前台用户 前缀区分
	TokenType                 = "Bearer"
	CacheModeRedis       int8 = 2
	BackendServerName         = "开源电商系统"
	MultiLogin                = true
	FrontendMultiLogin        = true
	GTokenExpireIn            = 10 * 24 * 60 * 60

	//文章相关
	ArticleIsAdmin = 1 //管理员发布
	ArticleIsUser  = 2 //用户发布

	//收藏相关
	CollectionTypeGoods   uint8 = 1
	CollectionTypeArticle uint8 = 2
)
