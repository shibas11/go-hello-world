/*
채널에 채널을 넣습니다.
즉 생산자가 소비자한테 채널을 통해 데이터를 보내는게 아니라, 소비자가 생산자한테 이제 나 준비됬으니깐, 니가 만든 것을 내가 준 채널을 통해서 보내줘~ 라는 의미입니다. 적극적인 소비자지요.
 */
package async

import (
	"fmt"
	"time"
)

func Test13() {
	cc := make(chan chan string)

	go consuming(cc)
	go producing(cc)

	time.Sleep(3 * time.Second)
}

func consuming(cc chan chan string) { // 적극적인 소비자
	resCh := make(chan string)

	cc <- resCh

	//value := <- resCh
	//fmt.Printf("응답: %v\n", value)
	for value := range resCh {
		fmt.Printf("응답: %v\n", value)
	}
}

func producing(cc chan chan string) {
	resCh := <-cc

	resCh <- "hello world!"
}
