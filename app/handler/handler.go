package handler

import (
	"github.com/verryp/cake-store-service/app/common"
	"github.com/verryp/cake-store-service/app/service"
)

type HandlerOption struct {
	*common.Option
	Service *service.Service
}
