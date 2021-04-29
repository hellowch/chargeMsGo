package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//定义中间件,用于输出请求信息
func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		//请求前输出信息
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("请求path:", path)
		fmt.Println("请求method:", method)

		//请求后输出的信息
		context.Next() //表示请求后执行
		fmt.Println("状态码:", context.Writer.Status())
	}
}
