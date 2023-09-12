package main

import (
	"main/internal/service/judge"
	"main/internal/service/judge/handle"
)

func main() {
	// 判题器
	go handle.Judger()

	if err := judge.Run(); err != nil {
		panic(err)
	}
}
