package service

import (
	"time"

	"github.com/verryp/cake-store-service/app/common"
	"github.com/verryp/cake-store-service/app/model"
	"github.com/verryp/cake-store-service/app/payload"
)

type ICakeService interface {
	Create(req payload.CakeCreateRequest) *common.Response
}

type cakeService struct {
	*ServiceOption
}

func NewCakeService(opt *ServiceOption) ICakeService {
	return &cakeService{opt}
}

func (s *cakeService) Create(req payload.CakeCreateRequest) *common.Response {
	c := &model.Cake{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
		CreatedAt:   time.Now(),
		// UpdatedAt:   req.UpdatedAt,
	}

	err := s.Repository.Cake.Create(c)
	if err != nil {
		s.Log.Err(err).Msg("failed to create cake")
		return common.APIResponse("error", common.ErrCreate.Error(), err)
	}

	return common.APIResponse("success", "successfully created cake", c)


}