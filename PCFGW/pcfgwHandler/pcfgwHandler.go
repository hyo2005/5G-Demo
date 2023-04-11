package pcfgwhandler

import (
	"github.com/gin-gonic/gin"
)

func HandlerRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"pcf": "2.2.2.2",
	})
}
