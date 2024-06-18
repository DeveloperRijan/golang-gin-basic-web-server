package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

// Home Page Handler
func HomeHandler(ctx *gin.Context) {
	//binding: query string to json struct
	var u User
	ctx.ShouldBind(&u)

	var response string
	response = fmt.Sprintf("Username: %s, email: %s", u.Name, u.Email)
	ctx.JSON(200, gin.H{
		"name":    u.Name,
		"email":   u.Email,
		"Welcome": response,
	})
}

func main() {
	r := gin.Default()

	r.GET("/", HomeHandler)

	r.Run()
}
