package chfgwhandler

import (
	"github.com/gin-gonic/gin"
)

func HandlerRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"chf": "3.3.3.3",
	})
}
