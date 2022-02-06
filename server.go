package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"

	"./lib/agentPool"
	"./lib/handleAgent"
	"./listeners/TcpListener"
)

func listagents() {
	pool := agentPool.GetAgents()
		fmt.Println("Slno.\tAgent Hash")
		for i := range pool {
			fmt.Println(i, "\t", pool[i])
		}
}

func main() {

	for {
		fmt.Print(">> ")
		scanner := bufio.NewScanner(os.Stdin)
		var cmd string
		if scanner.Scan() {
			cmd = scanner.Text()
		}

		command := strings.Split(cmd," ")
		switch command[0] {
		case "interact" :
			agenthash := command[1]
			c1 := agentPool.FetchAgent(agenthash)
			handleAgent.ClientCommunicate(c1,agenthash)
		case "remove":
			agenthash := command[1]
			agentPool.RemoveAgent(agenthash)
		case "listen":
			fmt.Print("Enter port : ")
			scanner := bufio.NewScanner(os.Stdin)
			var port string
			if scanner.Scan() {
				port = scanner.Text()
			}
			c := TcpListener.StartListener(port)
			_ = agentPool.AddToAgentPool(c)
		case "agents":
			listagents()

		default:
			fmt.Println("Wrong Command")
		}
	}
}

/*
func waitingConnection(l net.Listener) {
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = agentPool.AddToAgentPool(c)
	fmt.Println("Connected to agent at : ", c.RemoteAddr().String())
	return
}
*/
