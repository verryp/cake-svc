package server

import (
	"fmt"

	"github.com/verryp/cake-store-service/app/common"
	"github.com/verryp/cake-store-service/app/driver"
	"github.com/verryp/cake-store-service/app/handler"
	"github.com/verryp/cake-store-service/app/model"
	"github.com/verryp/cake-store-service/app/repository"
	"github.com/verryp/cake-store-service/app/service"
	"gopkg.in/gorp.v2"
)

func Start() {
	conf, err := common.NewConfig()
	if err != nil {
		panic(err)
	}

	logger := common.NewLogger(&conf.Log)

	opt := &common.Option{
		Config: conf,
		Log:    logger,
	}

	db, err := driver.NewMySQLDatabase(conf.DB)
	if err != nil {
		logger.Err(err).Msg("DB connection error")
		panic(err)
	}

	initDB(db)

	repos := wiringServerRepository(&repository.RepositoryOption{
		Option: opt,
		DB:     db,
	})

	services := wiringServerService(&service.ServiceOption{
		Option:     opt,
		Repository: repos,
	})

	handlers := &handler.HandlerOption{
		Option:  opt,
		Service: services,
	}

	svr := Router(handlers)

	addr := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)
	logger.Info().Msgf("server is serving at %s", addr)

	err = svr.Run(addr)
	if err != nil {
		logger.Err(err).Msgf("Server failed to serve at %s", addr)
	}
}

func wiringServerRepository(opt *repository.RepositoryOption) *repository.Repository {
	cake := repository.NewCakeRepo(opt)
	return &repository.Repository{
		Cake: cake,
	}
}

func wiringServerService(opt *service.ServiceOption) *service.Service {
	healthCheck := service.NewHealthCheckService(opt)
	cake := service.NewCakeService(opt)

	return &service.Service{
		HealthCheck: healthCheck,
		Cake: cake,
	}
}

func initDB(db *gorp.DbMap) {
	db.AddTableWithName(model.Cake{}, "cakes").SetKeys(true, "Id")
}
