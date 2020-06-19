/*
select 예제 6
  채널이 아닌, "채널의 값"을 넘겨줄 수 있음
  채널의값을 매개변수로 받는 함수는, 해당 채널에 값이 들어오지 않으면 call되지 않음
*/
package async

import (
	"fmt"
	"sync"
)

var ch6 chan string
var wg6 sync.WaitGroup = sync.WaitGroup{}

func Test06() {
	console := make(chan string, 1)
	ch6 = make(chan string, 1)

	wg6.Add(1)

	go func() {
		consuming6(<-console) // 채널이 아닌, "채널의 값"을 넘겨줄 수 있음
	}()

	go producing6(console)

	wg6.Wait()
}

func consuming6(prompt string) {
	fmt.Println("consuming 호출됨")

	select {
	case ch6 <- prompt:
		fmt.Println("이름 입력됨:", <-ch6)
	//case <-time.After(5 * time.Second):
	//	fmt.Println("타임 오버") // 함수 consuming6 자체가 호출되지 않으므로, 이 곳은 절대 실행되지 않음
	}

	wg6.Done()
}

func producing6(console chan string) {
	var name string
	fmt.Println("이름:")
	fmt.Scanln(&name)

	console <- name
}
