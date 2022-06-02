package HandleListeners

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nemesisdev2000/Nemesis/TeamServer/DataTypes"
	"github.com/nemesisdev2000/Nemesis/TeamServer/ListenerPool"
	"github.com/nemesisdev2000/Nemesis/TeamServer/Listeners/TcpListener"
)

type fn func(DataTypes.ListenerProfile)
type stop func(DataTypes.ListenerDetails)

type listenerID struct {
	ID string `json:"id"`
}

var listenerFunctions = map[string]fn{
	"TCP": TcpListener.StartListener,
}

var listenerStop = map[string]stop{
	"TCP": TcpListener.StopListener,
}

func HandleListener(c *gin.Context) {
	var listener DataTypes.ListenerProfile

	if err := c.BindJSON(&listener); err != nil {
		return
	}

	fmt.Println("Type : ", reflect.TypeOf(listener.Type))

	listenerType := strings.TrimSpace(listener.Type)

	go listenerFunctions[listenerType](listener)
	c.IndentedJSON(http.StatusOK, listener)
}

func ShowListeners(c *gin.Context) {
	var tcpListenerPool []ListenerPool.TcpListenerType
	tcpListenerPool = ListenerPool.ShowListeners()
	for _, a := range tcpListenerPool {
		c.IndentedJSON(http.StatusOK, a.ListenerID)
		c.IndentedJSON(http.StatusOK, a.TcpListener.Addr().String())
	}
	return
}

func StopListener(c *gin.Context) {
	var listenerDetails DataTypes.ListenerDetails

	if err := c.BindJSON(&listenerDetails); err != nil {
		return
	}

	listenerStop[listenerDetails.Type](listenerDetails)
}
