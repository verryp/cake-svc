package repository

import (
	"github.com/verryp/cake-store-service/app/model"
)

type CakeRepo interface {
	Create(cake *model.Cake) error
}

type cakeRepo struct {
	*RepositoryOption
}

func NewCakeRepo(opt *RepositoryOption) CakeRepo {
	return &cakeRepo{opt}
}

func (r *cakeRepo) List() {

}

func (r *cakeRepo) Create(c *model.Cake) error {
	// r.DB.Select(&test, "select id, name from JOIN  where $1 ", id)
	return r.DB.Insert(c)
}