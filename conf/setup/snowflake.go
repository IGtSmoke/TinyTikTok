package setup

import (
	"TinyTikTok/conf"
	"github.com/bwmarrin/snowflake"
)

var SnowflakeNode *snowflake.Node

func Snowflake() {
	SnowflakeNode, _ = snowflake.NewNode(conf.Conf.SnowflakeID)
}
