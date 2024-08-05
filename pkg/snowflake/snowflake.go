package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"math/rand"
	"time"
)

var node *sf.Node

func Init(startTime string, machineId int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-01", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineId)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
func IdNum() (randnum int64) {
	rand.Seed(time.Now().Unix())
	randnum = int64(rand.Intn(9999999999999999))
	return randnum
}
