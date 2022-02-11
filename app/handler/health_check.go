package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthCheckHandler struct {
	*HandlerOption
}

func NewHealthCheckHandler(opt *HandlerOption) *healthCheckHandler {
	return &healthCheckHandler{opt}
}

func (h *healthCheckHandler) Readiness(c *gin.Context) {
	ctx := c.Request.Context()

	err := h.Service.HealthCheck.HealthCheck(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := struct {
		HealthCheck string `json:"health_check"`
	}{HealthCheck: "up"}

	c.JSON(http.StatusOK, res)
}
