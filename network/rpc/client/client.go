package main

import (
	"fmt"
	"github.com/shibas11/go-hello-world/network/types"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:6000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	// 동기 호출
	fmt.Println("동기 호출")
	args := &types.Args{Operand1: 1, Operand2: 2}
	reply := new(types.Reply)
	err = client.Call("CalcService.Sum", args, reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("[sync]  reply.Result", reply.Result)

	// 비동기 호출
	fmt.Println("비동기 호출")
	args.Operand1 = 4
	args.Operand2 = 9
	sumCall := client.Go("CalcService.Sum", args, reply, nil) // 고루틴 호출
	<-sumCall.Done                                            // 채널 대기
	fmt.Println("[async] reply.Result", reply.Result)
}
