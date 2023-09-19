package chatgpt

import (
	"io"
	rpcChatGPT "main/api/chatgpt"
	"main/internal/common"
	"main/internal/common/build"
	"main/internal/common/ctxt"
	"main/rpc"

	"github.com/gin-gonic/gin"
)

type (
	Message struct {
		Role    string `json:"role" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	ChatRequest struct {
		Messages []*Message `json:"messages" binding:"required"`
	}
)

func Chat(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")
	defer c.Writer.Flush()

	// 解析参数
	req := new(ChatRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.SSEvent("error", common.CodeInvalidParams.Msg())
		return
	}

	// 构造消息
	var messages []*rpcChatGPT.Message
	builder := new(build.Builder)
	for _, m := range req.Messages {
		message := new(rpcChatGPT.Message)
		builder.Build(m, message)
		messages = append(messages, message)
	}
	if builder.Error() != nil {
		c.SSEvent("error", common.CodeServerBusy.Msg())
		return
	}

	// 调用Chat服务
	ctx, cancel := ctxt.WithTimeoutContext(60)
	defer cancel()
	stream, err := rpc.ChatGPTCli.Chat(ctx, &rpcChatGPT.ChatRequest{Messages: messages})
	if err != nil {
		panic(err)
	}

	// 接收消息流
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil || resp.GetStatusCode() != common.CodeSuccess.Code() {
			c.SSEvent("error", common.CodeServerBusy.Msg())
			break
		}
		c.SSEvent("msg", resp.GetContent())
		c.Writer.Flush()
	}
}
