package service

import "context"

type IHealthCheck interface {
	HealthCheck(ctx context.Context) error
}

type healthCheck struct {
	*ServiceOption
}

func NewHealthCheckService(opt *ServiceOption) IHealthCheck {
	return &healthCheck{opt}
}

func (s *healthCheck) HealthCheck(ctx context.Context) error {
	return nil
}
