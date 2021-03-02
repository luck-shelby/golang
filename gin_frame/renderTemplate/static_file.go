/**
 * @Author: luckfei
 * @Description: 我相信一切都是最好的安排
 * @File:  static_file
 * @Version: 1.15
 * @Date: 2021/1/2 上午3:00
 */

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 注释
func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(200,"welcome gin")
	})
	server := &http.Server{
		Addr: ":8099",
		Handler: router,
		ReadTimeout: time.Second * 10,
	}
	_ = server.ListenAndServe()
}
