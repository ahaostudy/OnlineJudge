package snowflake

import "main/services/user/config"

var node *Node

func init() {
	n, err := NewNode(int64(config.Config.Auth.Node))
	if err != nil {
		panic(err)
	}
	node = n
}

func Generate() ID {
	return node.Generate()
}