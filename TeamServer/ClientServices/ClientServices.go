package ClientServices

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []userDetails{}

func Login(c *gin.Context) {
	var user userDetails

	if err := c.BindJSON(&user); err != nil {
		fmt.Println("error")
		return
	}

	users = append(users, user)

	if user.Password == "Passw0rd" {
		c.IndentedJSON(http.StatusOK, user)
	} else {
		c.IndentedJSON(http.StatusBadRequest, "Wrong Password")
	}

}
