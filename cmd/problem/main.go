package main

import "main/internal/service/problem"

func main() {
	if err := problem.Run(); err != nil {
		panic(err)
	}
}
