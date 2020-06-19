package async

import (
	"fmt"
	"time"
)

func Test12() {
	c1 := make(chan int)

	go func() {
		for {
			i := <-c1
			fmt.Println("c1 :", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c1 <- 20
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case c1 <- 10:
			/* 사실 이 부분이 없어도 맨위 goroutine에서 모두 수신 가능함

			case i := <-c1: // 바로 위 case에서 채널에 넣은 10을, 여기서 출력하진 못함(동일 goroutine)
				fmt.Println("c2 : ", i)
			 */
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
