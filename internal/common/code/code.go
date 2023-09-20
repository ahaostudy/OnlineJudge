package code

// Code 响应状态码
type Code int64

const (
	CodeSuccess Code = 1000

	CodeInvalidParams     Code = 2001
	CodeUserExist         Code = 2002
	CodeUserNotExist      Code = 2003
	CodeInvalidPassword   Code = 2004
	CodeNotMatchPassword  Code = 2005
	CodeInvalidToken      Code = 2006
	CodeNotLogin          Code = 2007
	CodeInvalidCaptcha    Code = 2008
	CodeRecordNotFound    Code = 2009
	CodeSubmitNotFound    Code = 2010
	CodeNoRegistration    Code = 2011
	CodeContestNotStarted Code = 2012

	CodeServerBusy Code = 4001
)

var msg = map[Code]string{
	CodeSuccess: "success",

	CodeInvalidParams:     "请求参数错误",
	CodeUserExist:         "用户已存在",
	CodeUserNotExist:      "用户不存在",
	CodeInvalidPassword:   "用户名或密码错误",
	CodeNotMatchPassword:  "两次密码不一致",
	CodeInvalidToken:      "无效的Token",
	CodeNotLogin:          "用户未登录",
	CodeInvalidCaptcha:    "验证码错误",
	CodeRecordNotFound:    "记录不存在",
	CodeSubmitNotFound:    "提交不存在",
	CodeNoRegistration:    "未报名比赛",
	CodeContestNotStarted: "比赛未开始",

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
