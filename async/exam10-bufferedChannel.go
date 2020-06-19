package async

import "fmt"

func Test10() {
	/*
	  c := make(chan int)
	  c <- 1           // 수신 goroutine 없으므로 데드락
	  fmt.Println(<-c) // 이렇게 해도 데드락 (별도의 goroutine 없기 때문)
	*/

	ch := make(chan int, 1)
	ch <- 101          // buffered channel은 수신자가 없더라도 보낼 수 있음
	fmt.Println(<-ch)
}
