package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var base_api = "/api/core/test"

func main() {
	var r = gin.Default()
	controller(r)
	// 可以配置不同的端口
	r.Run(":8080")
}

func controller(t *gin.Engine) {
	var r = t.Group("/v1")
	r.GET(base_api+"/get1", get1)
	//r.Handle("get", base_api+"/get11", get1)
	r.GET(base_api+"/get2", get2)
	r.GET(base_api+"/json/ascii", someJson)
	r.GET(base_api+"/jsonP", jsonP)
	r.POST(base_api+"/login", login1)
}

func get1(c *gin.Context) {
	// {"message":"pong"}
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// 返回值
func get2(c *gin.Context) {
	// {"message":"pong"}

	result := map[string]string{}
	result["a"] = "ok1"
	result["b"] = "ok2"

	c.JSON(http.StatusOK, result)
}

func someJson(c *gin.Context) {
	data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
	}

	// {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}

// jsonP是能够解决跨域问题的
func jsonP(c *gin.Context) {
	data := map[string]interface{}{
		"foo": "bar",
	}

	// /JSONP?callback=x
	// 将输出：x({\"foo\":\"bar\"})
	c.JSONP(http.StatusOK, data)
}

type LoginForm struct {
	user string `form:"user"`
}

func login(c *gin.Context) {
	//var userInfo UserInfo
	//// 绑定OK
	//if c.ShouldBind(&userInfo) == nil {
	//	if userInfo.User == "zhou" && userInfo.Password == "simon" {
	//		c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
	//	} else {
	//		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	//	}
	//}

	var form LoginForm

	// 在这种情况下，将自动选择合适的绑定
	if c.Bind(&form) == nil {
		if form.user == "user" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}
}

func login1(c *gin.Context) {
	//var userInfo UserInfo
	//// 绑定OK
	//if c.ShouldBind(&userInfo) == nil {
	//	if userInfo.User == "zhou" && userInfo.Password == "simon" {
	//		c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
	//	} else {
	//		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	//	}
	//}

	//var form LoginForm
	//
	//// 在这种情况下，将自动选择合适的绑定
	//if c.Bind(&form) == nil {
	//	if form.user == "user" {
	//		c.JSON(200, gin.H{"status": "you are logged in"})
	//	} else {
	//		c.JSON(401, gin.H{"status": "unauthorized"})
	//	}
	//}

	logInReq := TestReq{}

	if errA := c.ShouldBind(&logInReq); errA == nil {
		fmt.Println(logInReq)
	}
}

type TestReq struct {
	Name string `json:"name" binding:"required"`
	Age  int8   `json:"age" binding:"required"`
}
