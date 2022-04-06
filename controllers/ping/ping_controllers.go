package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func String(c *gin.Context) {
	c.String(http.StatusOK, "Pong!")
	return
}

func Json(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong!",
	})
}
