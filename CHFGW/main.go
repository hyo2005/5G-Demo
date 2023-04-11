package main

import (
	chfgwhandler "CHFGW/chfgwHandler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/chf-gw", chfgwhandler.HandlerRequest)
	r.Run(":8084")
}
