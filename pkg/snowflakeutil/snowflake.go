package snowflakeutil

import (
	"github.com/bwmarrin/snowflake"
)

type Node struct {
	node *snowflake.Node
}

type Config struct {
	MachineID int64
}

func New(config *Config) *Node {
	node, err := snowflake.NewNode(config.MachineID)
	if err != nil {
		panic(err)
	}
	return &Node{
		node: node,
	}
}

func (n *Node) GenerateID() int64 {
	return n.node.Generate().Int64()
}
