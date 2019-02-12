package main

import (
	"testing"
	"fmt"
	"github.com/icloudland/rpc-scaffold/rpclient"
)

func TestClient1(t *testing.T) {

	connCfg := &rpcclient.ConnConfig{
		Host:         "127.0.0.1:8996",
		Endpoint:     "ws",
		User:         "admin",
		Pass:         "admin",
		HTTPPostMode: false,
		DisableTLS:   true,
	}

	c, err := rpcclient.New(connCfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Uptime())
	fmt.Println(c.Stop())

	select {

	}

}
