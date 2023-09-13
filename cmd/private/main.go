package main

import "main/internal/service/private"

func main() {
	if err := private.Run(); err != nil {
		panic(err)
	}
}
