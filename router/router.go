package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "gogofly/docs"
	global "gogofly/global/constants"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type IRegisterFn = func(public *gin.RouterGroup, private *gin.RouterGroup)

var (
	routers []IRegisterFn
)

func RegisterRoute(fn IRegisterFn) {
	if fn == nil {
		return
	}
	routers = append(routers, fn)
}

// 初始化 gin 框架路由信息
func InitRouter() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGTERM)
	defer cancel()

	r := gin.Default()
	// 对应的开发 API 的前缀
	public := r.Group("/api/v1/public")
	// 对应的需要认证的 API 的前缀
	auth := r.Group("/api/v1")

	// 初始化基础服务平台的 路由
	InitBasePlatformRoutes()

	// 注册自定义的校验器
	registerCusValidator()

	// 集成 swagger
	// 生成的 swagger 文档的访问交给项目来管理
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 注册路由
	for _, router := range routers {
		router(public, auth)
	}

	port := viper.GetString("server.port")
	if port == "" {
		port = "8900"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	go func() {
		global.Log.Infof("Start Listen :%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Error("Start Server Error: %s\n", err.Error())
			return
		}
		//
		global.Log.Errorf("Start Server Error: %s\n", port)
	}()

	// 阻塞程序运行，直到接收到关闭信号后触发了上下文中信道的关闭，
	<-ctx.Done()

	ctx, showdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer showdown()

	if err := server.Shutdown(ctx); err != nil {
		global.Log.Errorf("服务关闭异常: %s\n", err.Error())
		return
	}
	global.Log.Info("服务停止成功")
}

func InitBasePlatformRoutes() {
	InitUserRoutes()
	InitAbortRoutes()
}

// 自定义校验器
func registerCusValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("capitalized", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				if value != "" && strings.ToTitle(value) == value {
					return true
				}
			}
			return false
		})
	}
}
