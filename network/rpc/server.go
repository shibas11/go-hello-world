package rpc

import (
	"fmt"
	"github.com/shibas11/go-hello-world/network/types"
	"net"
	"net/rpc"
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

	rpc.Register(new(types.CalcService)) // types.CalcService 타입의 인스턴스를 RPC 서버에 등록

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

		go rpc.ServeConn(conn) // RPC를 처리하는 함수를 고루틴으로 실행
	}
}
