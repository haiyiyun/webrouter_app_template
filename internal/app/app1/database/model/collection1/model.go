package collection1

import (
	"project_name/internal/app/app1/database"

	"github.com/haiyiyun/cache"
	"github.com/haiyiyun/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type Model struct {
	*database.Database `json:"-" bson:"-" map:"-"`
}

func NewModel(mgo mongodb.Mongoer, col bson.M, cc *cache.Cache) *Model {
	obj := &Model{
		Database: database.NewDatabase(mgo, col, cc),
	}

	return obj
}
