package service

import (
	"github.com/verryp/cake-store-service/app/common"
	"github.com/verryp/cake-store-service/app/repository"
)

type ServiceOption struct {
	*common.Option
	Repository *repository.Repository
}

type Service struct {
	HealthCheck IHealthCheck
	Cake ICakeService
}
