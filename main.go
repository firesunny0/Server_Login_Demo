package main

import (
	"fmt"
	"net/http"
	"server_login/my_db"
	"server_login/user"

	"github.com/gin-gonic/gin"
)

func main() {
	my_db.InitDb()
	r := gin.Default()
	r.POST("/user/login", func(c *gin.Context) {
		jsonData := user.UserInfo{}
		c.BindJSON(&jsonData)
		fmt.Printf("log : %v", jsonData)
		err := jsonData.Login()
		if err != nil {
			c.String(http.StatusOK, "%v", err)
		} else {
			c.String(http.StatusOK, "Login success !")
		}
	})
	r.POST("/user/register", func(c *gin.Context) {
		jsonData := user.UserInfo{}
		c.BindJSON(&jsonData)
		fmt.Printf("log : %v", jsonData)
		err := jsonData.Register()
		if err != nil {
			c.String(http.StatusOK, "%v", err)
		} else {
			c.String(http.StatusOK, "Register success ! Welcome %v", jsonData.Name)
		}
	})
	r.Run(":8899")
}
