package async

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func Test11() {
	c := make(chan int)

	wg.Add(2)

	go producer(c)
	go consumer(c)

	wg.Wait()
}

func producer(c chan<- int) { // 보내기 전용 채널
	for i := 0; i <= 5000000; i++ {
		c <- i
	}

	c <- 10000000 // 채널에 값을 보냄
	//fmt.Println(<-c) // 채널에서 값을 꺼내면 컴파일 에러

	close(c) // 채널을 close
	wg.Done()
}

func consumer(c <-chan int) { // 받기 전용 채널
	// c <- 1        // 채널에 값을 보내면 컴파일 에러

	for i := range c {  // for range로 채널 수신, 채널이 close될 때까지
		if i%1000000 == 0 {
			fmt.Println(i)
		}
	}

	wg.Done()
}
