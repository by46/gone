package raft

import (
	"fmt"

	"net/http"

	"github.com/coreos/etcd/raft"
	"github.com/labstack/echo"
)

func Serve() http.Handler {
	storage := raft.NewMemoryStorage()
	c := &raft.Config{
		ID:              0x01,
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256,
	}
	peers := []raft.Peer{{ID: 0x01}}
	node := raft.StartNode(c, peers)
	fmt.Printf("Node %v\n", node.Status())

	app := echo.New()

	return app

}
