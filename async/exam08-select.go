/*
select 예제 8
  list의 item 수만큼만 for문 돌기
  select문으로 스케줄링한다

  done채널의 size만큼만 done채널에 값을 넣을 수 있으므로
  done채널의 처리량(size)가 꽉 차면
  "case done <- true" 구문에서 블럭된다
  고루틴 내부의 "<-done" 처리를 통해서 버퍼(done채널의 처리량)가 해소되면 다시 진행된다
*/
package async

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

var quitC = make(chan string)
var errC = make(chan error)

func Test08() {

	go testFn()

	select {
	case err := <-errC:
		fmt.Println("에러 발생", err)
	case <-quitC:
		fmt.Println("이상 없음")
	}
}

func testFn() {
	list := []string{"00", "11", "22", "33", "44", "55", "66", "77"}

	maxParallelFiles := 3
	done := make(chan bool, maxParallelFiles)

	wg8 := sync.WaitGroup{}
	for i, entry := range list {
		done <- true
		wg8.Add(1)

		go func(idx int, item string) {
			defer wg8.Done()

			err := readString(idx, item)
			if err != nil {
				errC <- err
				return
			}

			<-done

		}(i, entry)

	}

	wg8.Wait()
	quitC <- "finish"
}

func readString(idx int, str string) error {
	time.Sleep(1 * time.Second)

	if len(strings.TrimSpace(str)) == 0 {
		return errors.New("str is empty")
	}

	fmt.Printf(`%d:  str is "%s"
`, idx, str)
	return nil
}
