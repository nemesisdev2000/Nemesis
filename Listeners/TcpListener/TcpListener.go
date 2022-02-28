package TcpListener

import (
	"crypto/rand"
	"fmt"
	"net"
	"os"
	"reflect"
	"strconv"

	reuse "github.com/libp2p/go-reuseport"
)

func GenerateId() int {
	t, _ := rand.Prime(rand.Reader, 128)
	return int(t.Uint64())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type listenerConfig struct {
	Type     string
	Port     string
	ID       string
	conn     net.Conn
	listener net.Listener
}

var listenerPool []listenerConfig

func RemoveIndex(s []listenerConfig, i int) []listenerConfig {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func StartListener(port string) net.Conn {
	l, err := reuse.Listen("tcp4", ":"+port)
	fmt.Println("Listener type : ", reflect.TypeOf(l))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	HandleListener("TcpListener", port, c, l)
	fmt.Println(listenerPool)
	return c
}

func StopListener(id string) {
	for i, a := range listenerPool {
		fmt.Println("Index : ", i)
		if a.ID == id {
			fmt.Println("This ID  : ", a.ID)
			a.conn.Close()
			a.listener.Close()
			listenerPool = RemoveIndex(listenerPool, i)
			break
		}
	}
	return
}

func HandleListener(Type string, Port string, conn net.Conn, listen net.Listener) {
	id := strconv.Itoa(GenerateId())
	lconfig := listenerConfig{Type: Type, Port: Port, ID: id, conn: conn, listener: listen}
	listenerPool = append(listenerPool, lconfig)
	return
}
