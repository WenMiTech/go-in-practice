package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(log.Fields{
		"name": "age",
	}).Info("A log test")
	log.Info("log test 2")

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	r.Run()

}
