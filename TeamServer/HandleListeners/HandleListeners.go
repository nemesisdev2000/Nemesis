package HandleListeners

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nemesisdev2000/Nemesis/TeamServer/DataTypes"
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

	listenerFunctions[listenerType](listener)
	c.IndentedJSON(http.StatusOK, listener)
}
