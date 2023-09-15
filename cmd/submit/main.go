package main

import "main/internal/service/submit"

func main() {
	if err := submit.Run(); err != nil {
		panic(err)
	}
}
