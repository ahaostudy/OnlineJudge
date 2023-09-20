package main

import "main/internal/service/contest"

func main() {
	if err := contest.Run(); err != nil {
		panic(err)
	}
}