package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	gin.SetMode(gin.DebugMode)
}
func main() {
	r := gin.Default()
	//当做命名空间
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name": "test",
			})
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name": "test2",
			})
		})
	}
	r.Run(":3000")
}
