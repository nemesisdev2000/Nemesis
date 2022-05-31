package ListenerPool

import (
	"fmt"
	"net"
	"strings"

	"github.com/nemesisdev2000/Nemesis/TeamServer/CryptoFunctions"
)

type TcpListenerType struct {
	ListenerID  string
	TcpListener net.Listener
}

var TcpListenerPool []TcpListenerType

func remove(s []TcpListenerType, i int) []TcpListenerType {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

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

func DeleteListener(listenerID string) {
	for index, a := range TcpListenerPool {
		if strings.Contains(listenerID, a.ListenerID) {
			a.TcpListener.Close()
			TcpListenerPool = remove(TcpListenerPool, index)
			fmt.Println("Deleted Listener ID : ", a.ListenerID)
		}
	}
}
