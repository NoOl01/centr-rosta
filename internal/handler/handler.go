package handler

import (
	"centr_rosta/internal/handler/admin"
	"centr_rosta/internal/handler/auth"
	"centr_rosta/internal/handler/middleware"
)

type Handler struct {
	handlerAuth  auth.HandlerAuth
	handlerAdmin admin.HandlerAdmin
	middleware   middleware.Middleware
}

func NewHandler(handlerAuth auth.HandlerAuth, handlerAdmin admin.HandlerAdmin, middleware middleware.Middleware) *Handler {
	return &Handler{
		handlerAuth:  handlerAuth,
		handlerAdmin: handlerAdmin,
		middleware:   middleware,
	}
}
