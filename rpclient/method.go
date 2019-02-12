package rpcclient

import (
	"github.com/icloudland/rpc-scaffold/rpcjson"
)

func (c *Client) Uptime() (string, error) {
	cmd := rpcjson.NewUptimeCmd()
	result, err := receiveFuture(c.sendCmd(cmd))
	if err != nil {
		return "", err
	}

	return string(result[:]), nil
}


func (c *Client) Stop() (string, error) {
	cmd := rpcjson.NewStopCmd()
	result, err := receiveFuture(c.sendCmd(cmd))
	if err != nil {
		return "", err
	}

	return string(result[:]), nil
}