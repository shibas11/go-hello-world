package tcp

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	Port int
}

func NewServer(port int) *Server {
	s := new(Server)
	s.Port = port

	return s
}

func (s *Server) Serve() error {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(s.Port))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go requestHandler(conn) // 패킷을 처리할 함수를 고루틴으로 실행
	}
}

func requestHandler(conn net.Conn) {
	defer conn.Close()

	data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스

	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Received:", string(data[:n]))

		_, err = conn.Write(data[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
