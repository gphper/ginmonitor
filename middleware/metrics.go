/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-08-04 22:39:47
 */
package middleware

import (
	"github.com/gin-gonic/gin"
)

func Metrics(c *gin.Context) {
	c.Next()
}
