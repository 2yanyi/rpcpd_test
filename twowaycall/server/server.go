// Source code file, created by Developer@YANYINGSONG.

package main

import (
	"library/generic/errcause"
	"r/tool"
	"rpcpd"
	"rpcpd/seqx"
)

const serverAddr = "127.0.0.1:12099"

var power *rpcpd.Power

func runServer() {
	seqx.Init(2)

	var powerErr error
	power, powerErr = rpcpd.Listen(serverAddr)
	if powerErr != nil {
		panic(powerErr)
	}
	defer power.Listener.Close()

	for {
		conn, err := power.Accept()
		if err != nil {
			panic(err)
		}

		go connectionLive(conn)

		// TEST
		// go tool.CallTest(nil, "S", power, conn, 0)
		go tool.CallTestLoop("S", power, conn)
	}
}

func connectionLive(conn *rpcpd.Conn) {
	defer errcause.Recover()
	if err := power.ConnectionProcessor(conn); err != nil {
		panic(err)
	}
}

func main() {
	// rpcpd.DebugOn()
	rpcpd.AddFunction(tool.TestCallFunctionName, tool.Foo)
	runServer()
}
