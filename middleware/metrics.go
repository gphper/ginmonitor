/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-04 22:39:47
 */
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	gapro "github.com/gphper/ginmonitor/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

func Metrics(c *gin.Context) {
	startTime := time.Now()
	gapro.HttpRequestCounter.With(prometheus.Labels{"method": c.Request.Method, "path": c.FullPath()}).Inc()
	c.Next()
	endTime := time.Now()
	gapro.HttpRquestTime.With(prometheus.Labels{"method": c.Request.Method, "path": c.FullPath()}).Set(float64(endTime.Sub(startTime)))
}
