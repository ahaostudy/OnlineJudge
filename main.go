package main

import (
	"log"
	submit "main/kitex_gen/submit/submitservice"
)

func main() {
	svr := submit.NewServer(new(SubmitServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
