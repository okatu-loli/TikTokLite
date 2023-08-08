package dyerr

const (
	//状态码
	OK           uint32 = 200
	NO_REQUEST   uint32 = 400
	WRONG_SERVER uint32 = 500
	//状态信息
	SUCCESS string = "成功执行！！！！"

	//服务错误码
	SERVER_COMMON_ERROR uint32 = 100001
	//请求参数错误
	REUQEST_PARAM_ERROR uint32 = 100002
	//token过期
	TOKEN_EXPIRE_ERROR uint32 = 100003
	//token生成错误
	TOKEN_GENERATE_ERROR uint32 = 100004
	//数据库错误
	DB_ERROR uint32 = 100005
	//数据库更新错误
	DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006
)

var (
	ErrGenerateTokenError       = NewErrMsg("生成token失败")
	ErrUsernamePwdError         = NewErrMsg("账号或密码不正确")
	ErrUserNoExistsError        = NewErrMsg("用户不存在")
	ErrUserAlreadyRegisterError = NewErrMsg("该用户名已被注册")
	ErrDBerror                  = NewErrMsg("数据库错误")
	ErrServerCommonError        = NewErrMsg("服务器请求错误")
	ErrRequestParamError        = NewErrMsg("请求参数错误")
	ErrTokenExpireError         = NewErrMsg("Token已经过期")
	ErrUpdateAffectedZeroError  = NewErrMsg("数据库更新错误")
)
