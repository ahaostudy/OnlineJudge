package main

import (
	"main/internal/service/judge"
)

func main() {
	if err := judge.Run(); err != nil {
		panic(err)
	}
}
