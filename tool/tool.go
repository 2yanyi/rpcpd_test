// Source code file, created by Developer@YANYINGSONG.

package tool

import (
	"errors"
	"fmt"
	"rpcpd"
	"rpcpd/seqx"
	"sync"
	"time"
)

const TestCallFunctionName = "golang.foo.method"

func Foo(data []byte) ([]byte, error) {
	time.Sleep(time.Second)
	return data, nil
}

func CallTestLoop(tag string, power *rpcpd.Power, conn *rpcpd.Conn) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go CallTest(wg, tag, power, conn, i)
	}
}

func CallTest(wg *sync.WaitGroup, tag string, power *rpcpd.Power, conn *rpcpd.Conn, i int) {
	if wg != nil {
		defer wg.Done()
	}
	time.Sleep(1 * time.Second)
	// fmt.Println("发起调用：服务端调客户端")

	packet := []byte(fmt.Sprintf("%d %s:%d", i, tag, seqx.X.NextID()))
	resp, err := power.Call(conn, []byte(TestCallFunctionName), packet)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}

	// 结果正确性验证
	if string(packet) != string(resp) {
		fmt.Printf("%s != %s\n", packet, resp)
		panic(errors.New("data bad"))
	}

	fmt.Printf("调用结束： 发(%s) 收(%s)\n", packet, resp)
}
