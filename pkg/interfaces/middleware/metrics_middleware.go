package middleware

import "github.com/gin-gonic/gin"

func MetricMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
