package main

import "main/internal/service/chatgpt"

func main() {
	if err := chatgpt.Run(); err != nil {
		panic(err)
	}
}
