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
	//url参数
	//精确匹配
	//匹配 /pv/12  /pv/ or /pv 都不会匹配
	r.GET("/pv/:date", func(c *gin.Context) {
		date := c.Param("date")
		c.String(http.StatusOK, "%s", date)
	})
	///pv2/:date/*hour
	//匹配 /pv2/12/12
	r.GET("/pv2/:date/*hour", func(c *gin.Context) {
		hour := c.Param("hour")
		c.String(http.StatusOK, "time is %s", hour)
	})

	//查询参数
	// /pv3?date=12
	r.GET("/pv3", func(c *gin.Context) {
		date := c.Query("date")
		hour := c.DefaultQuery("hour", "0") //hour default value is 0
		c.String(http.StatusOK, "date is %s hour is %s", date, hour)
	})
	r.Run()
}
