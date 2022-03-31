package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingString(c *gin.Context) {
	c.String(http.StatusOK, "Pong!")
	return
}

func PingJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong!",
	})
}
