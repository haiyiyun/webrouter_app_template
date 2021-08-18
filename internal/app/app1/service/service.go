package service

import (
	"project_name/internal/app"

	"github.com/haiyiyun/cache"
)

type Service struct {
	*app.Service
	*Config
	*cache.Cache
}

func NewService(c *Config, cc *cache.Cache) *Service {
	return &Service{
		Config: c,
		Cache:  cc,
	}
}
