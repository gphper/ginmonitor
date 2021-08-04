/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-04 21:52:27
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
