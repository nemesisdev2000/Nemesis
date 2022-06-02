package DataTypes

import (
	"fmt"
)

type ListenerProfile struct {
	Type string `json:"type"`
	Port string `json:"port"`
	Host string `json:"host"`
}

type ListenerDetails struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func Nothing() {
	fmt.Println("Nothing here")
}
