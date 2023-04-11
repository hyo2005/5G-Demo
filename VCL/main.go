package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hyo2005/demo-5g/initializer"
	"github.com/hyo2005/demo-5g/logic"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/logic-create", logic.CreateSubscriber)
	r.GET("/logic-view/:id", logic.ViewSubscriber)
	r.PUT("/logic-update", logic.UpdateSubscriber)
	r.DELETE("/logic-delete/:id", logic.DeleteSubscriber)
	r.Run(":8888")
}
