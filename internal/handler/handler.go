package handler

import (
	"centr_rosta/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Router(r *gin.Engine)
}

type handler struct {
	Service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handler{Service: service}
}
