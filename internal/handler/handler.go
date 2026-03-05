package handler

import (
	"centr_rosta/internal/handler/admin"
	"centr_rosta/internal/handler/auth"
	"centr_rosta/internal/handler/middleware"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Router(r *gin.Engine)
}

type handler struct {
	handlerAuth  auth.HandlerAuth
	handlerAdmin admin.HandlerAdmin
	middleware   middleware.Middleware
}

func NewHandler(handlerAuth auth.HandlerAuth, handlerAdmin admin.HandlerAdmin, middleware middleware.Middleware) Handler {
	return &handler{
		handlerAuth:  handlerAuth,
		handlerAdmin: handlerAdmin,
		middleware:   middleware,
	}
}
