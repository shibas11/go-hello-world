package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	go func(conn net.Conn) {
		data := make([]byte, 4096)

		for {
			n, err := conn.Read(data)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(data[:n]))

			time.Sleep(1 * time.Second)
		}
	}(client)

	go func(conn net.Conn) {
		i := 0
		for {
			s := "Hello " + strconv.Itoa(i)

			_, err := conn.Write([]byte(s))
			if err != nil {
				fmt.Println(err)
				return
			}

			i++
			time.Sleep(1 * time.Second)
		}
	}(client)

	fmt.Scanln()
}
