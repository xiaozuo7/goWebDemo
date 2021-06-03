package errmsg

const (
	Success = 200
	Error   = 500

	ErrorTokenExists  = 1001
	ErrorTokenType    = 1002
	ErrorTokenExpired = 1003

	ErrorUserNameExists   = 2001
	ErrorUserNotExists    = 2002
	ErrorPasswordInvalid  = 2003
	ErrorPermissionDenied = 2004
)

var codeMsg = map[int]string{
	Success: "OK",
	Error:   "服务器错误",

	ErrorTokenExists:  "token不存在",
	ErrorTokenType:    "token格式错误",
	ErrorTokenExpired: "token已过期",

	ErrorUserNameExists:   "用户名已存在",
	ErrorUserNotExists:    "用户不存在",
	ErrorPasswordInvalid:  "密码错误",
	ErrorPermissionDenied: "没有权限登录",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
