package test

import (
	"net"
	"testing"
)

type client struct {
	conn *net.TCPConn
}

func Test_Agent(t *testing.T)  {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:10001")
	if err != nil {
		t.Fatalf("tcp resolve err:, %v", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	var msg string
	msg = "666"

	for i := 0; i < 3; i ++ {
		_, err := conn.Write([]byte(msg))
		if err != nil {
			t.Logf("send error: %v", err)
			continue
		}
	}
}

