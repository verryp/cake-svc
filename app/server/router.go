package server

import (
	"github.com/gin-gonic/gin"
	"github.com/verryp/cake-store-service/app/handler"
)

func Router(opt *handler.HandlerOption) *gin.Engine {
	hcHandler := handler.NewHealthCheckHandler(opt)
	cHandler := handler.NewCakeHandler(opt)

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	health := r.Group("health")
	health.GET("/readiness", hcHandler.Readiness)

	cakes := r.Group("cakes")
	cakes.POST("/", cHandler.CreateCake)
	// cakes.GET("/", ) // list
	// cakes.GET("/{id}", ) // detail by id
	// cakes.PUT("/{id}", ) // update by id
	// cakes.DELETE("/{id}", ) // delete by id


	return r
}
