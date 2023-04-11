package main

import (
	ipallocatehandler "IPAllocate/ipAllocateHandler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ip-allocate", ipallocatehandler.HandlerRequest)
	r.Run(":8081")
}
