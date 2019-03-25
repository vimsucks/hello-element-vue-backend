package util

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

var snowflakeNode *snowflake.Node

func init() {
	var err error
	snowflakeNode, err = snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}
}

func NewUID() snowflake.ID {
	return snowflakeNode.Generate()
}
