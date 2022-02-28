package C2Server

import (
	"net"

	"github.com/nemesisdev2000/Nemesis/Listeners/TcpListener"
)

func Call(port string, f func(string) net.Conn) net.Conn {
	return f(port)
}

func Listen(port string, ltype string) net.Conn {
	m := map[string]func(string) net.Conn{
		"TcpListener": TcpListener.StartListener,
	}

	c := Call(port, m[ltype])
	return c
}

func Stop(id string) {
	TcpListener.StopListener(id)
	return
}
