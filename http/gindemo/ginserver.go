package main

import (
	"math/rand"
	"go.uber.org/zap"
	"time"
	"github.com/gin-gonic/gin"
)


const keyRequestID = "requestId"


func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	r.Use(func(c *gin.Context) {
		s := time.Now()

		c.Next()

		logger.Info("incoming request", 
		zap.String("path", c.Request.URL.Path), 
		zap.Int("status", c.Writer.Status()),
		zap.Duration("elapsed", time.Now().Sub(s)),
	)
	}, func(c *gin.Context) {
		c.Set(keyRequestID, rand.Int())
		c.Next()
	})


	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(keyRequestID); exists {
			h[keyRequestID] = rid
		}
		c.JSON(200, h)
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}