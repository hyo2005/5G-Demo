package ipallocatehandler

import (
	"github.com/gin-gonic/gin"
)

func HandlerRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"ipallocate": "0.0.0.0",
	})
}
