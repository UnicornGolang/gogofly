package cmd

import (
	"fmt"
	"gogofly/config"
	global "gogofly/global/constants"
	"gogofly/router"
	"gogofly/utils"
)

func Start() {
	var initErr error
	fmt.Println("=======================系统启动=====================")

	//==================================================
	// 初始化系统配置
	config.InitConfig()

	//==================================================
	// 初始化日志组件
	global.Log = config.InitLogger()

	//==================================================
	// 初始化数据配置
	db, err := config.InitDb()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	//==================================================
	// 初始化 redis 配置

	rdb, err := config.InitRedis()
	global.RDB = rdb
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	global.RDB.Set("alex", "more")
	fmt.Println(global.RDB.Get("alex"))

	//==================================================
	// 检测启动时是否有错误
	if initErr != nil {
		if global.Log != nil {
			global.Log.Error(initErr.Error())
		}
		panic(initErr.Error())
	}
	test_token()
	//==================================================
	// 初始化系统路由组件, 开启之后将会进入阻塞接收请求
	//
	router.InitRouter()
}

// 清理工作
func Clean() {
	global.Log.Info("=======================系统清理=====================")
}

func test_token() {
	token, err := utils.GenerateToken(12138, "alex")
	if err != nil {
		panic("err" + err.Error())
	}
	fmt.Println(token)
	valid := utils.InValid(token)
	fmt.Println(valid)
	claims, err := utils.ParseToken(token)
	if err != nil {
		panic("err" + err.Error())
	}
	fmt.Println(claims.Id, claims.Subject, claims.ExpiresAt)
}
