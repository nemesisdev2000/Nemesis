package TcpListener

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/nemesisdev2000/Nemesis/TeamServer/DataTypes"
	"github.com/nemesisdev2000/Nemesis/TeamServer/ListenerPool"
)

func StartListener(listener DataTypes.ListenerProfile) {
	host := listener.Host
	port := listener.Port
	l, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	fmt.Println("Listenening on " + host + ":" + port)

	go handleRequest(l)
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	fmt.Println("Connection established : ", conn.RemoteAddr().String())
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	temp := strings.TrimSpace(string(data))
	fmt.Println("Callback received from : ", temp)

	return
}

func StopListener(listenerDetails DataTypes.ListenerDetails) {
	ListenerPool.DeleteListener(listenerDetails.ID)
}

func handleRequest(l net.Listener) {
	fmt.Println("Adding Listener")
	ListenerPool.AddListener(l)
}
