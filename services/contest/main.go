package main

import (
	"log"
	contest "main/kitex_gen/contest/contestservice"
)

func main() {
	svr := contest.NewServer(new(ContestServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
