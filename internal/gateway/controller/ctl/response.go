package ctl

import "main/internal/common"

// 基础响应体
type Response struct {
	StatusCode common.Code   `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func (r *Response) CodeOf(code common.Code) Response {
	if r == nil {
		r = new(Response)
	}
	r.StatusCode = code
	r.StatusMsg = code.Msg()
	return *r
}

func (r *Response) Success() {
	r.CodeOf(common.CodeSuccess)
}
