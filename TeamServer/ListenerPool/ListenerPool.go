package ListenerPool

import (
	"fmt"
	"net"

	"github.com/nemesisdev2000/Nemesis/TeamServer/CryptoFunctions"
)

type TcpListenerType struct {
	ListenerID  string
	TcpListener net.Listener
}

var TcpListenerPool []TcpListenerType

func AddListener(l net.Listener) {
	//adding the tcplistener to listener pool
	var listener TcpListenerType
	listener.ListenerID = CryptoFunctions.GetMD5Hash(CryptoFunctions.GenerateRandomString())
	listener.TcpListener = l
	TcpListenerPool = append(TcpListenerPool, listener)
}

func ShowListeners() []TcpListenerType {
	fmt.Println("Showing Listeners ")
	return TcpListenerPool
}

func StopListener(c net.Conn, l net.Listener) {
	c.Close()
	l.Close()
}
