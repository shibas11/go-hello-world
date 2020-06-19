/*
select 예제 2
  생산자가 결과를 줄 때까지 기다리는 방식의 코드
    - process 함수가 종료되면, 그 결과를 받아서 종료(return문)
 */

package async

import (
	"fmt"
	"time"
)

func Test02() {
	ch := make(chan string)
	go process(ch)

	n := 1
	for {
		fmt.Println("\n----------", n, "----------")
		time.Sleep(1 * time.Second)

		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return // 종료
		default:
			fmt.Println("no value received")
		}

		fmt.Println("do something")
		n = n + 1
	}
}

func process(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "process successful"
}
