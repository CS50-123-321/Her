package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {
	r := gin.Default()
	r.GET("/info", func(c *gin.Context) {
		aiRes, err := GetAiResponse()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"response": aiRes,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
