package upfgwhandler

import (
	"github.com/gin-gonic/gin"
)

func HandlerRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"upf": "1.1.1.1",
	})
}
