package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verryp/cake-store-service/app/payload"
)

type cakeHandler struct {
	*HandlerOption
}

func NewCakeHandler(opt *HandlerOption) *cakeHandler {
	return &cakeHandler{opt}
}

func (h *cakeHandler) CreateCake(c *gin.Context) {
	var input payload.CakeCreateRequest

	err := c.ShouldBindJSON(&input)
	h.Log.Info().Msg(input.Title)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	// input.CreatedAt = time.Now()
	res := h.Service.Cake.Create(input)

	c.JSON(http.StatusOK, res)
	return
}