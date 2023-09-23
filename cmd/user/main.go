package main

import "main/internal/service/user"

func main() {
	if err := user.Run(); err != nil {
		panic(err)
	}
}
