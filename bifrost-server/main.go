package main

import (
	"github.com/aloks98/bifrost/bifrost-server/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db.Migrate()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":10056")
}
