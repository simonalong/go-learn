package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lunny/log"
	"github.com/simonalong/mikilin-go"
	toolsLog "github.com/simonalong/tools/log"
	"net/http"
	"time"
)

var base_api = "/api/core/test"

func main() {
	var r = gin.Default()
	// 配置耗时
	r.Use(timeoutMiddleware(time.Second * 1))
	//controller(r)
	r.GET("get1", get1)
	r.GET("get2", timedHandler2)
	r.POST("post1", post1)
	r.POST("handle", handle)
	r.GET("short", timedHandler(time.Second*5))

	toolsLog.LogRouters(r)
	// 可以配置不同的端口
	r.Run(":8080")
}

var watchMap map[string]chan string = make(map[string]chan string)

func timeoutMiddleware2(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.JSON(http.StatusRequestTimeout, "request timeout")
				c.Abort()
			}
			cancel()
		}()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

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

func timedHandler2(c *gin.Context) {
	ctx := c.Request.Context()

	doneChan := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		doneChan <- "sdf"
	}()

	select {
	case <-ctx.Done():
		fmt.Println("done")
		return
	case res := <-doneChan:
		c.JSON(http.StatusOK, res)
	}
}

func timedHandler(duration time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		doneChan := make(chan string)
		go func() {
			time.Sleep(duration)
		}()

		watchMap["key"] = doneChan

		select {

		case <-ctx.Done():
			return

		case res := <-doneChan:
			c.JSON(http.StatusOK, res)
		}
	}
}

func controller(t *gin.Engine) {
	var r = t.Group("/v1")
	r.GET(base_api+"/get1/:id", get1)
	//r.Handle("get", base_api+"/get11", get1)
	//r.GET(base_api+"/get2", get2)
	//r.GET(base_api+"/get3", get3)
	//r.GET(base_api+"/json/ascii", someJson)
	//r.GET(base_api+"/jsonP", jsonP)
	//r.POST(base_api+"/login", login1)
}

func get1(c *gin.Context) {
	ip, _ := c.RemoteIP()
	time.Sleep(10 * time.Second)
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "success",
		"message": "成功",
		"data":    c.ClientIP() + ", " + ip.String(),
	})
}

func post1(c *gin.Context) {
	// {"message":"pong"}
	dataReq := DataReq{}
	err := c.BindJSON(&dataReq)
	if err != nil {
		log.Errorf("err：%v", err.Error())
		return
	}

	checkResult, errMsg := mikilin.Check(dataReq)
	if !checkResult {
		log.Warnf(errMsg)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "success",
		"message": "成功",
		"data":    dataReq,
	})
	//ip, _ := c.RemoteIP()
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"code":    "success",
	//	"message": "成功",
	//	"data":    c.ClientIP() + ", " + ip.String(),
	//})

}

func handle(c *gin.Context) {
	chanStr := watchMap["key"]
	chanStr <- "hahahahah"

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "success",
		"message": "成功",
		"data":    "keyil",
	})
}

type DataReq struct {
	Age int `json:"age" match:"range=[0, 12]" errMsg:"age不合法，合法值：0~12，包含边界，当前：#current"`
}

//
////// 返回值
//func get2(c *gin.Context) {
//	// {"message":"pong"}
//
//	result := map[string]string{}
//	result["a"] = "ok1"
//	result["b"] = "ok2"
//
//	c.JSON(http.StatusOK, result)
//}
//
//// 返回值
//func get3(c *gin.Context) {
//	// {"message":"pong"}
//}
//
//func someJson(c *gin.Context) {
//	data := map[string]interface{}{
//		"lang": "GO语言",
//		"tag":  "<br>",
//	}
//
//	// {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
//	c.AsciiJSON(http.StatusOK, data)
//}
//
//// jsonP是能够解决跨域问题的
//func jsonP(c *gin.Context) {
//	data := map[string]interface{}{
//		"foo": "bar",
//	}
//
//	// /JSONP?callback=x
//	// 将输出：x({\"foo\":\"bar\"})
//	c.JSONP(http.StatusOK, data)
//}
//
//type LoginForm struct {
//	user string `form:"user"`
//}
//
//func login(c *gin.Context) {
//	//var userInfo UserInfo
//	//// 绑定OK
//	//if c.ShouldBind(&userInfo) == nil {
//	//	if userInfo.User == "zhou" && userInfo.Password == "simon" {
//	//		c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
//	//	} else {
//	//		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
//	//	}
//	//}
//
//	var form LoginForm
//
//	// 在这种情况下，将自动选择合适的绑定
//	if c.Bind(&form) == nil {
//		if form.user == "user" {
//			c.JSON(200, gin.H{"status": "you are logged in"})
//		} else {
//			c.JSON(401, gin.H{"status": "unauthorized"})
//		}
//	}
//}
//
//func login1(c *gin.Context) {
//	//var userInfo UserInfo
//	//// 绑定OK
//	//if c.ShouldBind(&userInfo) == nil {
//	//	if userInfo.User == "zhou" && userInfo.Password == "simon" {
//	//		c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
//	//	} else {
//	//		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
//	//	}
//	//}
//
//	//var form LoginForm
//	//
//	//// 在这种情况下，将自动选择合适的绑定
//	//if c.Bind(&form) == nil {
//	//	if form.user == "user" {
//	//		c.JSON(200, gin.H{"status": "you are logged in"})
//	//	} else {
//	//		c.JSON(401, gin.H{"status": "unauthorized"})
//	//	}
//	//}
//
//	logInReq := TestReq{}
//
//	if errA := c.ShouldBind(&logInReq); errA == nil {
//		fmt.Println(logInReq)
//	}
//}
//
//type TestReq struct {
//	Name string `json:"name" binding:"required"`
//	Age  int8   `json:"age" binding:"required"`
//}
