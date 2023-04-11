package main

import (
	"HTTPGW/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/create-subscriber", service.HandleCreateSub)
	r.GET("/view-subscriber", service.HandleViewSub)
	r.PUT("/update-subscriber", service.HandleUpdateSub)
	r.DELETE("/delete-subscriber", service.HandleDeleteSub)
	r.Run(":8080")
}
