package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 为了解决服务端的支持跨域，引入了中间件
// go get -u github.com/gin-contrib/cors
// 在其中添加跨域允许的站点，方法，响应头，cookie 信息等
func Cors() gin.HandlerFunc {
	cfg := cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "token"},
		AllowCredentials: true,
	}
	return cors.New(cfg)
}
