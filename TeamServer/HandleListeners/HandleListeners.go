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

func HandleListener(c *gin.Context) {
	var listener DataTypes.ListenerProfile

	if err := c.BindJSON(&listener); err != nil {
		return
	}

	fmt.Println("Type : ", reflect.TypeOf(listener.Type))

	listenerType := strings.TrimSpace(listener.Type)

	listenerFunctions := map[string]fn{
		"TCP": TcpListener.StartListener,
	}

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
