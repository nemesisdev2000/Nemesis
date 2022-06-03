package AgentHandler

import (
	"github.com/nemesisdev2000/Nemesis/TeamServer/DataTypes"
)

type AgentProfile struct {
	Type     string
	Port     string
	BindHost string
	AgentID  string
}

func fn(AgentProfile)

func RegisterAgent(listenerDetails DataTypes.ListenerDetails) {
}
