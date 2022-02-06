package handleAgent

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"net"
	"os"
)

func ClientCommunicate(c net.Conn, agenthash string) {
	var cmdex string
	for {
		fmt.Println("Interacting with agent at : ",c.RemoteAddr().String())
		fmt.Print(agenthash," >> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			cmdex = scanner.Text()
		}
		if cmdex == "quit" {
			enc := b64.StdEncoding.EncodeToString([]byte(cmdex))
			cmdex = enc
			fmt.Fprint(c, cmdex+"\n")
			c.Close()
			break
		}
		enc := b64.StdEncoding.EncodeToString([]byte(cmdex))
		cmdex = enc

		fmt.Println("Encoded Message ", cmdex)
		fmt.Fprint(c, cmdex+"\n")
		netData, err := bufio.NewReader(c).ReadString('\n')
		output, _ := b64.StdEncoding.DecodeString(netData)
		if err != nil {
			fmt.Println("Error")
			c.Close()
			os.Exit(0)
		}
		fmt.Println(string(output))
	}
	return
}
