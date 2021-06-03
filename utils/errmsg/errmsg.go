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

	ErrorFileSizeNotAllow = 3001
)

var codeMsg = map[int]string{
	Success: "OK",
	Error:   "服务器错误",

	ErrorTokenExists:  "Token不存在",
	ErrorTokenType:    "Token格式错误",
	ErrorTokenExpired: "Token已过期",

	ErrorUserNameExists:   "用户名已存在",
	ErrorUserNotExists:    "用户不存在",
	ErrorPasswordInvalid:  "密码错误",
	ErrorPermissionDenied: "没有权限登录",

	ErrorFileSizeNotAllow: "上传文件超过最大可上传范围",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
