package test

import "main/config"

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
}
