package middleware

import (
	"gin_skeleton/g"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"net/http"
	"sync"
)

// ip限流
func IpRateLimit(rate int) gin.HandlerFunc {

	g.RateLimits = &sync.Map{}
	return func(c *gin.Context) {
		ip := c.ClientIP()
		_, ok := g.RateLimits.Load(ip)
		if !ok {
			//限流单ip
			g.RateLimits.Store(ip, ratelimit.New(rate))
		}
		// 取令牌
		rateLimitOne, ok := g.RateLimits.Load(ip)
		if !ok {
			c.AbortWithStatusJSON(http.StatusOK, nil)
		}
		rateLimitOne.(ratelimit.Limiter).Take()
		c.Next()
	}
}
