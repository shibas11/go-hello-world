package async

import (
	"fmt"
	"time"
)

func Test14() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("hello")
		case <-done:
			return
		}
	}
}
