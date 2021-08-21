package app1

import (
	"context"
	"flag"

	"project_name/internal/app/app1/database/schema"
	"project_name/internal/app/app1/service/base"
	app1Service1 "project_name/internal/app/app1/service/service1"

	"github.com/haiyiyun/cache"
	"github.com/haiyiyun/config"
	"github.com/haiyiyun/mongodb"
	"github.com/haiyiyun/webrouter"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	baseConfFile := flag.String("config.app1.base", "../config/app1/base.conf", "base config file")
	var baseConf base.Config
	config.Files(*baseConfFile).Load(&baseConf)

	baseCache := cache.New(baseConf.CacheDefaultExpiration.Duration, baseConf.CacheCleanupInterval.Duration)
	baseDB := mongodb.NewMongoPool("", baseConf.MongoDatabaseName, 100, options.Client().ApplyURI(baseConf.MongoDNS))
	webrouter.SetCloser(func() { baseDB.Disconnect(context.TODO()) })

	baseDB.M().InitCollection(schema.Collection1)

	baseService := base.NewService(&baseConf, baseCache, baseDB)

	if baseConf.WebRouter {
		//Init Begin
		app1Service1Service := app1Service1.NewService(baseService)
		//Init End

		//Go Begin
		//Go End

		//Register Begin
		webrouter.Register(baseConf.WebRouterRootPath+"", app1Service1Service)
		//Register End
	}
}
