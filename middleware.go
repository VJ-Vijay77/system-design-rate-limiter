package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"fmt"
)

const (
	requestLimit   = 10          // Max 10 requests
	windowDuration = time.Minute // Per 1 minute window
)

type clientData struct {
	Count     int
	LastReset time.Time
}

var (
	store = make(map[string]*clientData)
	mu    sync.Mutex
)

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		mu.Lock()
		defer mu.Unlock()

		data, exists := store[ip]
		if !exists {
			store[ip] = &clientData{Count: 1, LastReset: now}
			c.Next()
			return
		}

		if now.Sub(data.LastReset) > windowDuration {
			data.Count = 1
			data.LastReset = now
			c.Next()
			return
		}

		if data.Count > requestLimit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded " + fmt.Sprint(data.Count),
			})
			c.Abort()
			return
		}

		data.Count++
		c.Next()

	}
}
