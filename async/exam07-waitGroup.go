/*
WaitGroup: 고루틴이 끝날 때를 기다림
  Add와 Done의 개수가 일치해야 함. 일치하지 않으면 에러 발생
    fatal error: all goroutines are asleep - deadlock!
 */
package async

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	i  int
	mu sync.Mutex
}

func (c *counter) increment() {
	c.mu.Lock()
	c.i += 1
	fmt.Println("i 증가됨:", c.i)
	c.mu.Unlock()
}

func (c *counter) display() {
	fmt.Println(c.i)
}

func Test07() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		// goroutine 밖에서 설정해야 함
		wg.Add(1)     // WaitGroup이 기다리는 고루틴 1개 추가
		//wg.Add(2)          // WaitGroup이 기다리는 고루틴 n개 추가

		go func() {
			// 당연히 waitGroup.Done()은 고루틴 안에서 호출해야 함
			defer wg.Done()      // 고루틴 종료 시 1개 Done() 처리
			//defer wg.Add(-2)   // 고루틴 종료 시 n개 Done() 처리

			c.increment()
		}()
	}

	wg.Wait() // 모든 고루틴 끝날 때까지 대기

	c.display()

}
