package test

import (
	"fmt"
	"main/server/utils/snowflake"
	"testing"
)

func TestSnowflake(t *testing.T) {
	node, _ := snowflake.NewNode(1)

	for i := 0; i < 10; i++ {
		id := node.Generate().Int64()
		fmt.Println(id)
	}
}

func TestSnowflakeGenerate(t *testing.T) {
	ch := make(chan int64)
	for i := 0; i < 100; i++ {
		go func() {
			ch <- snowflake.Generate().Int64()
		}()
	}

	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}
