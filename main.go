package main

import (
	"account-manager-web/amwhttp"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.RedirectFixedPath = true
	r.Static("/_env", "./public/_env")
	r.Static("/amis", "./public/amis")
	r.Static("/pages", "./public/pages")
	r.Static("/xterm", "./public/xterm")

	r.GET("/", func(c *gin.Context) {
		c.File("public/index.html")
	})
	r.POST("/_api", amwhttp.ApiFunc)
	r.Run(":8085")
}
