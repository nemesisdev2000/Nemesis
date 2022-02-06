package TcpListener

import(
	"fmt"
	"net"
	"os"
	reuse "github.com/libp2p/go-reuseport"
)

var l net.Listener

func StartListener(port string) net.Conn {
	l, err := reuse.Listen("tcp4", ":"+port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Connected to agent at : ", c.RemoteAddr().String())
	return c
}

func StopListener() {
	l.Close()
	return
}
