package main

import (
	"testing"
	"time"
)

func TestRpcServer_Start(t *testing.T) {

	c := &rpcserverConfig{
		StartupTime:  time.Now().Unix(),
		DisableTLS:   true,
		RPCListeners: []string{"0.0.0.0:8996"},
		RPCUser:      "admin",
		RPCPass:      "admin",
	}

	rpcServer, err := newRPCServer(c)
	if err != nil {
		panic(err)
	}

	rpcServer.Start()

	select {}
}
