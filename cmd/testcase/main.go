package main

import "main/internal/service/testcase"

func main() {
	if err := testcase.Run(); err != nil {
		panic(err)
	}
}
