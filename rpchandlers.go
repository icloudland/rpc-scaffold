package main

import (
	"github.com/icloudland/rpc-scaffold/rpcjson"
	"time"
)

type commandHandler func(*rpcServer, interface{}, <-chan struct{}) (interface{}, error)

// rpcHandlers maps RPC command strings to appropriate handler functions.
// This is set by init because help references rpcHandlers and thus causes
// a dependency loop.
var rpcHandlers map[string]commandHandler
var rpcHandlersBeforeInit = map[string]commandHandler{
	"stop":    handleStop,
	"uptime":  handleUptime,
	"version": handleVersion,
}

// Commands that are currently unimplemented, but should ultimately be.
var rpcUnimplemented = map[string]struct{}{
}

// Commands that are available to a limited user
var rpcLimited = map[string]struct{}{
	// Websockets AND HTTP/S commands
	"help": {},
}


// handleStop implements the stop command.
func handleStop(s *rpcServer, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	select {
	case s.requestProcessShutdown <- struct{}{}:
	default:
	}
	return "server stopping.", nil
}

// handleUptime implements the uptime command.
func handleUptime(s *rpcServer, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	return time.Now().Unix() - s.cfg.StartupTime, nil
}

// handleVersion implements the version command.
//
// NOTE: This is a btcsuite extension ported from github.com/decred/dcrd.
func handleVersion(s *rpcServer, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	result := map[string]rpcjson.VersionResult{
		"btcdjsonrpcapi": {
			VersionString: jsonrpcSemverString,
			Major:         jsonrpcSemverMajor,
			Minor:         jsonrpcSemverMinor,
			Patch:         jsonrpcSemverPatch,
		},
	}
	return result, nil
}