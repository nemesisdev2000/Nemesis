package TcpListener

import (
	"fmt"
	"net"

	"github.com/nemesisdev2000/Nemesis/TeamServer/DataTypes"
)

func StartListener(listener DataTypes.ListenerProfile) {
	host := listener.Host
	port := listener.Port
	l, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	defer l.Close()

	fmt.Println("Listenening on " + host + ":" + port)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error", err.Error())
			return
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	reqlen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	conn.Write([]byte("Message received"))
	fmt.Println("Req len : ", reqlen)
	conn.Close()
}
