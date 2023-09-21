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
	CodeNotRegistred      Code = 2011
	CodeAlreadyRegistered Code = 2012
	CodeContestNotStarted Code = 2013
	CodeContestHasStarted Code = 2014
	CodeContestNotExist   Code = 2015
	CodeContestNotOngoing Code = 2016

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
	CodeNotRegistred:      "未报名比赛",
	CodeAlreadyRegistered: "用户已报名",
	CodeContestNotStarted: "比赛未开始",
	CodeContestHasStarted: "比赛已经开始",
	CodeContestNotExist:   "比赛不存在",
	CodeContestNotOngoing: "比赛未开始或已结束",

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
