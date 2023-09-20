package test

import (
	"fmt"
	"io"
	"main/api/chatgpt"
	"main/internal/common/code"
	"main/internal/common/ctxt"
	"main/rpc"
	"testing"
)

func TestChatGPT(t *testing.T) {
	err := rpc.InitGRPCClients()
	if err != nil {
		panic(err)
	}
	defer rpc.CloseGPRCClients()

	ctx, cancel := ctxt.WithTimeoutContext(30)
	defer cancel()

	stream, err := rpc.ChatGPTCli.Chat(ctx, &rpcChatGPT.ChatRequest{
		Messages: []*rpcChatGPT.Message{
			{
				Role:    "user",
				Content: "用python写个小游戏",
			},
		},
	})
	if err != nil {
		panic(err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if resp.GetStatusCode() != code.CodeSuccess.Code() {
			panic("error")
		}
		fmt.Print(resp.GetContent())
	}
}
