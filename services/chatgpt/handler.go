package main

import (
	chatgpt "main/kitex_gen/chatgpt"
)

// ChatGPTServiceImpl implements the last service interface defined in the IDL.
type ChatGPTServiceImpl struct{}

func (s *ChatGPTServiceImpl) Chat(req *chatgpt.ChatRequest, stream chatgpt.ChatGPTService_ChatServer) (err error) {
	println("Chat called")
	return
}
