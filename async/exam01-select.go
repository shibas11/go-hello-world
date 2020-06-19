/*
select 예제 1
  "case문의 채널"에 값이 들어올 때까지 select문에서 블록됨
  단, default문을 넣어주면, 블록되지 않고 순회됨
 */
package async

import (
	"fmt"
	"time"
)

func Test01() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			c1 <- "one"
		}
	}()

	go func() {
		for {
			time.Sleep(4 * time.Second)
			c2 <- "two"
		}
	}()

	n := 1
	for {
		fmt.Println("------------------ select ------------------", n)

		select {
		case msg1 := <-c1:
			fmt.Println(time.Now(), "received", msg1)
		case msg2 := <-c2:
			fmt.Println(time.Now(), "received", msg2)
			//default:
			//	fmt.Println("나는 default")
			//	time.Sleep(1 * time.Second)
		}
		n = n + 1
	}
}
