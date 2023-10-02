// Code generated by Kitex v0.7.2. DO NOT EDIT.

package chatgptservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	chatgpt "main/kitex_gen/chatgpt"
)

func serviceInfo() *kitex.ServiceInfo {
	return chatGPTServiceServiceInfo
}

var chatGPTServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ChatGPTService"
	handlerType := (*chatgpt.ChatGPTService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Chat": kitex.NewMethodInfo(chatHandler, newChatArgs, newChatResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "chatgpt",
		"ServiceFilePath": ``,
	}
	extra["streaming"] = true
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func chatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st := arg.(*streaming.Args).Stream
	stream := &chatGPTServiceChatServer{st}
	req := new(chatgpt.ChatRequest)
	if err := st.RecvMsg(req); err != nil {
		return err
	}
	return handler.(chatgpt.ChatGPTService).Chat(req, stream)
}

type chatGPTServiceChatClient struct {
	streaming.Stream
}

func (x *chatGPTServiceChatClient) Recv() (*chatgpt.ChatResponse, error) {
	m := new(chatgpt.ChatResponse)
	return m, x.Stream.RecvMsg(m)
}

type chatGPTServiceChatServer struct {
	streaming.Stream
}

func (x *chatGPTServiceChatServer) Send(m *chatgpt.ChatResponse) error {
	return x.Stream.SendMsg(m)
}

func newChatArgs() interface{} {
	return &ChatArgs{}
}

func newChatResult() interface{} {
	return &ChatResult{}
}

type ChatArgs struct {
	Req *chatgpt.ChatRequest
}

func (p *ChatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(chatgpt.ChatRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ChatArgs) Unmarshal(in []byte) error {
	msg := new(chatgpt.ChatRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChatArgs_Req_DEFAULT *chatgpt.ChatRequest

func (p *ChatArgs) GetReq() *chatgpt.ChatRequest {
	if !p.IsSetReq() {
		return ChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChatArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChatArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChatResult struct {
	Success *chatgpt.ChatResponse
}

var ChatResult_Success_DEFAULT *chatgpt.ChatResponse

func (p *ChatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(chatgpt.ChatResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ChatResult) Unmarshal(in []byte) error {
	msg := new(chatgpt.ChatResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChatResult) GetSuccess() *chatgpt.ChatResponse {
	if !p.IsSetSuccess() {
		return ChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*chatgpt.ChatResponse)
}

func (p *ChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChatResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Chat(ctx context.Context, req *chatgpt.ChatRequest) (ChatGPTService_ChatClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "Chat", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &chatGPTServiceChatClient{res.Stream}
	if err := stream.Stream.SendMsg(req); err != nil {
		return nil, err
	}
	if err := stream.Stream.Close(); err != nil {
		return nil, err
	}
	return stream, nil
}