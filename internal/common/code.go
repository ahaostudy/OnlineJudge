package common

// Code 响应状态码
type Code int64

const (
	CodeSuccess Code = 1000

	CodeInvalidParams    Code = 2001
	CodeUserExist        Code = 2002
	CodeUserNotExist     Code = 2003
	CodeInvalidPassword  Code = 2004
	CodeNotMatchPassword Code = 2005
	CodeInvalidCaptcha   Code = 2006
	CodeRecordNotFound   Code = 2007
	CodeSubmitNotFound   Code = 2008

	CodeInvalidToken Code = 3001
	CodeNotLogin     Code = 3002

	CodeServerBusy Code = 4001
)

var msg = map[Code]string{
	CodeSuccess: "success",

	CodeInvalidParams:    "请求参数错误",
	CodeUserExist:        "用户已存在",
	CodeUserNotExist:     "用户不存在",
	CodeInvalidPassword:  "用户名或密码错误",
	CodeNotMatchPassword: "两次密码不一致",
	CodeInvalidCaptcha:   "验证码错误",
	CodeRecordNotFound:   "记录不存在",
	CodeSubmitNotFound:   "判题提交不存在或结果已被读取",

	CodeInvalidToken: "无效的Token",
	CodeNotLogin:     "用户未登录",

	CodeServerBusy: "服务繁忙",
}

func (code Code) Code() int64 {
	return int64(code)
}

// Msg 获取响应消息
func (code Code) Msg() string {
	if m, ok := msg[code]; ok {
		return m
	}
	return msg[CodeServerBusy]
}
