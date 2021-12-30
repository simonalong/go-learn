package test

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

// timeout middleware wraps the request context with a timeout
func timeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				fmt.Println("超时1")
				c.Abort()
			}
			fmt.Println("超时2")
			cancel()
		}()
		fmt.Println("超时3")
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func timedHandler(duration time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		type responseData struct {
			status int
			body   map[string]interface{}
		}

		doneChan := make(chan responseData)

		go func() {
			time.Sleep(duration)
			doneChan <- responseData{
				status: 200,
				body:   gin.H{"hello": "world"},
			}
		}()

		select {
		case <-ctx.Done():
			return
		case res := <-doneChan:
			c.JSON(res.status, res.body)
		}
	}
}

func TestSub1(t *testing.T) {
	// create new gin without any middleware
	engine := gin.New()

	// add timeout middleware with 2 second duration
	engine.Use(timeoutMiddleware(time.Second * 2))

	// create a handler that will last 1 seconds
	engine.GET("/short", timedHandler(time.Second))

	// create a route that will last 5 seconds
	engine.GET("/long", timedHandler(time.Second*5))

	// run the server
	engine.Run(":8080")
}
