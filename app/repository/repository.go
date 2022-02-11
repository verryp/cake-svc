package repository

import (
	"github.com/verryp/cake-store-service/app/common"
	"gopkg.in/gorp.v2"
)

type RepositoryOption struct {
	*common.Option
	DB *gorp.DbMap
}

type Repository struct {
	Cake CakeRepo
}
