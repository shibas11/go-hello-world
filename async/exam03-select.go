/*
select 예제 3
  랜덤선택(이지만 완전 랜덤은 아님)
 */
package async

import (
	"fmt"
	"time"
)

func Test03() {
	output1 := make(chan string)
	output2 := make(chan string)

	go serve1(output1)
	go serve2(output2)

	time.Sleep(1 * time.Second) // sleep이 없으면 항상 나중 것인 "serve2"가 출력되네?

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

func serve1(ch chan<- string) {
	ch <- "serve1"
}

func serve2(ch chan<- string) {
	ch <- "serve2"
}
