package main

import (
	pcfgwhandler "PCFGW/pcfgwHandler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/pcf-gw", pcfgwhandler.HandlerRequest)
	r.Run(":8083")
}
