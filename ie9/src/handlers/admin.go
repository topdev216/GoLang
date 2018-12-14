package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func AuthRequired(c *gin.Context) {
	log.Println("需要登陆一下子")
}

func AdminLogin(c *gin.Context) {
	log.Println("login page")
}

func AdminLogout(c *gin.Context) {
	log.Println("logout page")
}
