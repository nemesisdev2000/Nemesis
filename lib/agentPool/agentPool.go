package agentPool

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"fmt"
	"net"
)

var pool = make(map[string]net.Conn)

func GenerateAgentHash(c net.Conn) string {
	agenthash := c.RemoteAddr().String()
	agenthash = b64.StdEncoding.EncodeToString([]byte(agenthash))
	hasher := sha1.New()
	hasher.Write([]byte(agenthash))
	agenthash = b64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return agenthash
}

func AddToAgentPool(c net.Conn) string{
	agenthash := GenerateAgentHash(c)
	fmt.Println("Agent hash : ",agenthash)
	pool[agenthash] = c
	return agenthash
}

func GetAgents() map[string]net.Conn {
	return pool
}

func FetchAgent(agenthash string) net.Conn {
	return pool[agenthash]
}

func RemoveAgent(agenthash string) {
	delete(pool,agenthash)
	fmt.Println("Deleted agent : ",agenthash)
}
