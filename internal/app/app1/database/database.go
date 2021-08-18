package database

import (
	"project_name/internal/app"

	"github.com/haiyiyun/cache"
	"github.com/haiyiyun/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type Database struct {
	*app.Database `json:"-" bson:"-" map:"-"`
	*cache.Cache  `json:"-" bson:"-" map:"-"`
}

func NewDatabase(mgo mongodb.Mongoer, col bson.M, cc *cache.Cache) *Database {
	mdl := &Database{
		Database: app.NewDatabase(mgo, col),
		Cache:    cc,
	}

	return mdl
}
