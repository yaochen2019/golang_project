package cmd

import (
	"fmt"
	"github.com/yaochen2019/gogofly/conf"
	"github.com/yaochen2019/gogofly/global"
	"github.com/yaochen2019/gogofly/router"
	"github.com/yaochen2019/gogofly/util"
)

func Start() {
	var initErr error
	//初始化系统配置文件
	conf.InitConfig()
	//初始化日志组件
	global.Logger = conf.InitLogger()

	//初始化数据库连接
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = util.AppendError(initErr, err)

	}

	//初始化Redis
	rdClient, err := conf.InitRedis()
	global.RedisClient = rdClient
	if err != nil {
		initErr = util.AppendError(initErr, err)
	}
	global.RedisClient.Set("username", "zs")
	fmt.Println(global.RedisClient.Get("username"))
	//判断初始化过程中是否存在错误
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())

	}
	//初始化系统路由
	router.InitRouter()
}

func Clean() {
	fmt.Print("===========clearn===========")

}
