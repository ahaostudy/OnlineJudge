package chatgpt

import (
	"encoding/json"
	"io"
	"log"
	"strings"

	"main/common/code"
	chatgpt "main/kitex_gen/chatgpt"
	"main/services/chatgpt/config"
	"main/services/chatgpt/pkg/request"
	"main/services/chatgpt/pkg/result"
)

// ChatGPTServiceImpl implements the last service interface defined in the IDL.
type ChatGPTServiceImpl struct{}

func (s *ChatGPTServiceImpl) Chat(req *chatgpt.ChatRequest, stream chatgpt.ChatGPTService_ChatServer) (err error) {
	b, _ := json.Marshal(req.GetMessages())
	log.Println("request chat:", string(b))
	// url
	baseUrl := strings.TrimSuffix(config.Config.Openai.BaseUrl, "/")
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
	r.SetHeader("Authorization", "Bearer "+config.Config.Openai.ApiKey)
	r.SetData(map[string]interface{}{
		"model":    config.Config.Openai.Model,
		"messages": messages,
		"stream":   true,
	})

	// response
	resp, err := r.POST()
	if err != nil {
		stream.Send(&chatgpt.ChatResponse{
			StatusCode: code.CodeServerBusy.Code(),
		})
		return nil
	}
	defer resp.Body.Close()
	log.Printf("openai response status: %v\n", resp.Status)

	// 流式读取
	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			stream.Send(&chatgpt.ChatResponse{
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
				stream.Send(&chatgpt.ChatResponse{
					StatusCode: code.CodeServerBusy.Code(),
				})
				break
			}

			// 响应
			stream.Send(&chatgpt.ChatResponse{
				StatusCode: code.CodeSuccess.Code(),
				Content:    s.Choices[0].Delta.Content,
			})
		}
	}

	return nil
}
