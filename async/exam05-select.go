/*
select 예제 5
  select-case문에서 바로! 채널에 값 할당 가능
 */
package async

import (
	"fmt"
	"sync"
	"time"
)

var wg5 sync.WaitGroup = sync.WaitGroup{}

func Test05() {
	ch := make(chan string, 1)
	prompt := "HAMA"

	wg5.Add(1)

	select {
	case ch <- prompt: // select-case문에서 바로! 채널에 값 할당 가능
		fmt.Println("이름:", <-ch)
		wg5.Done()
	case <-time.After(time.Second):
		fmt.Println("타임 오버")
		wg5.Done()
	}

	wg5.Wait()
}
