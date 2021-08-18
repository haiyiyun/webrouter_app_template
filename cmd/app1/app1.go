package main

import (
	"context"
	"flag"

	"project_name/internal/app/app1/database/schema"
	"project_name/internal/app/app1/service"
	app1ServiceService1 "project_name/internal/app/app1/service/service1"

	"github.com/haiyiyun/cache"
	"github.com/haiyiyun/config"
	"github.com/haiyiyun/mongodb"
	"github.com/haiyiyun/webrouter"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	app1ConfFile := flag.String("config.app1", "../config/app1/app1.conf", "app1 config file")
	var app1Conf service.Config
	config.Files(*app1ConfFile).Load(&app1Conf)

	app1Cache := cache.New(app1Conf.CacheDefaultExpiration.Duration, app1Conf.CacheCleanupInterval.Duration)
	app1DB := mongodb.NewMongoPool("", app1Conf.MongoDatabaseName, 100, options.Client().ApplyURI(app1Conf.MongoDNS))
	webrouter.SetCloser(func() { app1DB.Disconnect(context.TODO()) })

	app1DB.M().InitCollection(schema.Collection1)
	app1Service := service.NewService(&app1Conf, app1Cache)

	//Init Begin
	app1ServiceService1Service := app1ServiceService1.NewService(app1Service, app1DB)
	//Init End

	//Go Begin
	//Go End

	//Register Begin
	webrouter.Register("/", app1ServiceService1Service)
	//Register End
}
