package main

import (
	upfgwhandler "UPFGW/upfgwHandler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/upf-gw", upfgwhandler.HandlerRequest)
	r.Run(":8082")
}
