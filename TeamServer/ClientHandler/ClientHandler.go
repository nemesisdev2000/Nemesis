package ClientHandler

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"strings"
)

func HandleClients(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		decodedstring, _ := base64.StdEncoding.DecodeString(temp)
		fmt.Println("Received string : %s\n", string(decodedstring))
		if temp == "STOP" {
			break
		}

		c.Write([]byte(string("Received\n")))
	}
	c.Close()
}
