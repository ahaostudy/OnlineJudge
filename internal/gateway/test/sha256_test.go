package test

import (
	"fmt"
	"main/config"
	"main/internal/gateway/utils/sha256"
	"testing"
)

func TestSha256(t *testing.T) {
	fmt.Println("salt:", config.ConfServer.Salt)
	pwd := sha256.Encrypt("123456")
	fmt.Println(pwd)
}