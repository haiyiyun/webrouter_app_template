package service1

import (
	"project_name/internal/app/app1/service/base"
)

type Service struct {
	*base.Service
}

func NewService(s *base.Service) *Service {
	return &Service{
		Service: s,
	}
}
