/*
select 예제 4
  시간 내에 입력이 되는지 기다림
 */
package async

import (
	"fmt"
	"sync"
	"time"
)

var wg4 sync.WaitGroup = sync.WaitGroup{}

func Test04() {
	ch := make(chan string)

	wg4.Add(1)

	go consuming4(ch)
	go producing4(ch)

	wg4.Wait()
}

func consuming4(ch chan string) {
	select {
	case name := <-ch:
		fmt.Println("입력받은 이름:", name)
	case <-time.After(5 * time.Second):
		fmt.Println("타임 오버")
	}

	wg4.Done()
}

func producing4(ch chan string) {
	var name string
	fmt.Println("이름:")
	fmt.Scanln(&name)

	ch <- name
}
