package service1

import (
	"project_name/internal/app/app1/service"

	"github.com/haiyiyun/mongodb"
)

type Service struct {
	*service.Service
	M mongodb.Mongoer
}

func NewService(s *service.Service, m mongodb.Mongoer) *Service {
	return &Service{
		Service: s,
		M:       m,
	}
}
