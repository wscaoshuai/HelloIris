package datasource

import (
	"HelloIris/18-CmsProject/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions/sessiondb/redis"
)

/**
 * 返回Redis实例
 */
func NewRedis() *redis.Database {

	var database *redis.Database

	//项目配置
	cmsConfig := config.InitConfig()
	if cmsConfig != nil {
		iris.New().Logger().Info(" hello ")
		rd := cmsConfig.Redis
		iris.New().Logger().Info(rd)
		//database = redis.New(service.Config{
		//	Network:     rd.NetWork,
		//	Addr:        rd.Addr + ":" + rd.Port,
		//	IdleTimeout: time.Duration(24) * time.Hour,
		//	Password:    "",
		//	Prefix:      rd.Prefix,
		//})
		database = redis.New(redis.Config{
			Network:  "tcp",
			Addr:     "127.0.0.1:6379",
			Password: "",
			Database: "",
			//MaxIdle:     0,
			MaxActive: 10,
			//IdleTimeout: service.DefaultRedisIdleTimeout,
			Prefix: "",
		})
	} else {
		iris.New().Logger().Info(" hello  error ")
	}
	return database
}
