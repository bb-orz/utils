package xrandom

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/bwmarrin/snowflake"
	uuid "github.com/satori/go.uuid"
)

/**
生成分布式唯一ID: 整数,自增,且不连续
*/
func GenUniqueId() (snowflake.ID,error) {
	// Create a new Node with a Node number of 10000
	var nodeNumber int64 = 900
	node, err := snowflake.NewNode(nodeNumber + int64(randomdata.Number(99)))
	if err != nil {
		return -1,err
	}

	// Generate a snowflake ID.
	id := node.Generate()

	return id,nil
}

/**
生成uuid
*/
func GenUUID() string {
	uuids, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return uuids.String()
}
