package main

//
//
//import (
//	"github.com/gin-gonic/gin"
//	"log"
//)
//
//func main() {
//	r := gin.Default()
//	r.Static("/assets", "./assets")
//	//r.SetHTMLTemplate(html)
//
//	r.GET("/", func(c *gin.Context) {
//		if pusher := c.Writer.Pusher(); pusher != nil {
//			// 使用 pusher.Push() 做服务器推送
//			if err := pusher.Push("/assets/app.js", nil); err != nil {
//				log.Printf("Failed to push: %v", err)
//			}
//		}
//		c.HTML(200, "https", gin.H{
//			"status": "success",
//		})
//	})
//
//	// 监听并在 https://127.0.0.1:8080 上启动服务
//	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
//}
