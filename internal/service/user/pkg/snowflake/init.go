package snowflake

import "main/config"

var node *Node

func init() {
	n, err := NewNode(int64(config.ConfAuth.Node))
	if err != nil {
		panic(err)
	}
	node = n
}

func Generate() ID {
	return node.Generate()
}
