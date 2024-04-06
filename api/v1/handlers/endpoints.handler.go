package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tylorkolbeck/go-cookbook/internal/service"
)

type EndpointsHandler struct {
	service *service.EndpointsService
}

func NewEndpointsHandler(service *service.EndpointsService) *EndpointsHandler {
	return &EndpointsHandler{
		service: service,
	}
}

func (h *EndpointsHandler) ListEndpoints(c *gin.Context) {
	endpoints := h.service.Get()
	c.JSON(http.StatusOK, endpoints)
}
