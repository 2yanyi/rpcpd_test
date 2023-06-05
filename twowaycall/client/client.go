// Source code file, created by Developer@YANYINGSONG.

package main

import (
	"library/generic/errcause"
	"r/tool"
	"rpcpd"
	"rpcpd/seqx"
)

const serverAddr = "127.0.0.1:12099"

var client *rpcpd.Power

func runClient() func() {
	seqx.Init(1)

	var err error
	client, err = rpcpd.Connect(serverAddr)
	if err != nil {
		panic(err)
	}
	return func() {
		client.Client.Channel.Close()
	}
}

func connectionLive() {
	defer errcause.Recover()
	if err := client.ConnectionProcessor(nil); err != nil {
		panic(err)
	}
}

func main() {
	// rpcpd.DebugOn()
	defer runClient()()
	rpcpd.AddFunction(tool.TestCallFunctionName, tool.Foo)

	go tool.CallTest(nil, "C", client, nil, 0)
	// go tool.CallTestLoop("C", client, nil)

	// 保持长连接
	for i := 0; i < rpcpd.DefaultTaskNum; i++ {
		go connectionLive()
	}
	connectionLive()
}
