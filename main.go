package main

import "time"

func main() {

	interrupt := interruptListener()
	defer rpcsLog.Infof("Shutdown complete")

	c1 := rpcserverConfig{
		StartupTime:  time.Now().Unix(),
		DisableTLS:   true,
		RPCListeners: []string{"0.0.0.0:8996"},
		RPCUser:      "admin",
		RPCPass:      "admin",
	}

	rpcs, err := newRPCServer(&c1)
	if err != nil {
		rpcsLog.Errorf("Unable to start server on %v: %v",
			c1.Listeners, err)
		return
	}

	defer func() {
		rpcsLog.Infof("Gracefully shutting down the server...")
		rpcs.Stop()
	}()

	rpcs.Start()

	// Wait until the interrupt signal is received from an OS signal
	<-interrupt
}
