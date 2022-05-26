package DataTypes

import (
	"fmt"
)

type ListenerProfile struct {
	Type string `json:"type"`
	Port string `json:"port"`
	Host string `json:"host"`
}

func Nothing() {
	fmt.Println("Nothing here")
}
