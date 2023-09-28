package chatgpt

import (
	"encoding/json"
	"io"
	"log"
	"strings"

	rpcChatGPT "main/api/chatgpt"
	"main/config"
	"main/internal/common/code"
	"main/internal/service/chatgpt/pkg/request"
	"main/internal/service/chatgpt/pkg/result"
)

func (ChatGPTServer) Chat(req *rpcChatGPT.ChatRequest, stream rpcChatGPT.ChatGPTService_ChatServer) error {
	b, _ := json.Marshal(req.GetMessages())
	log.Println("request chat:", string(b))
	// url
	baseUrl := strings.TrimSuffix(config.ConfChatGPT.Openai.BaseUrl, "/")
	url := baseUrl + "/v1/chat/completions"

	// messages
	var messages []map[string]string
	for _, m := range req.Messages {
		messages = append(messages, map[string]string{
			"role":    m.GetRole(),
			"content": m.GetContent(),
		})
	}

	// request
	r := request.NewRequest(url)
	r.SetHeader("Content-Type", "application/json")
	r.SetHeader("Authorization", "Bearer "+config.ConfChatGPT.Openai.ApiKey)
	r.SetData(map[string]interface{}{
		"model":    config.ConfChatGPT.Openai.Model,
		"messages": messages,
		"stream":   true,
	})

	// response
	resp, err := r.POST()
	if err != nil {
		stream.Send(&rpcChatGPT.ChatResponse{
			StatusCode: code.CodeServerBusy.Code(),
		})
		return nil
	}
	defer resp.Body.Close()

	// 流式读取
	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			stream.Send(&rpcChatGPT.ChatResponse{
				StatusCode: code.CodeServerBusy.Code(),
			})
			break
		}
		// TODO: 此处使用`data: `分隔不太优雅
		lines := strings.Split(string(buf[:n]), "data: ")
		for _, line := range lines {
			// 处理
			line = strings.TrimSpace(strings.Trim(line, "\n"))
			if len(line) == 0 {
				continue
			}
			if line == "[DONE]" {
				break
			}

			// 解析
			s := new(result.ChatStream)
			if err := json.Unmarshal([]byte(strings.Trim(line, "\n")), s); err != nil || len(s.Choices) == 0 {
				stream.Send(&rpcChatGPT.ChatResponse{
					StatusCode: code.CodeServerBusy.Code(),
				})
				break
			}

			// 响应
			stream.Send(&rpcChatGPT.ChatResponse{
				StatusCode: code.CodeSuccess.Code(),
				Content:    s.Choices[0].Delta.Content,
			})
		}
	}

	return nil
}
