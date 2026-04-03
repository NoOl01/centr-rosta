package handler

import (
	"centr_rosta/internal/handler/admin"
	"centr_rosta/internal/handler/admin/admin_user"
	"centr_rosta/internal/handler/auth"
	"centr_rosta/internal/handler/lesson"
	"centr_rosta/internal/handler/middleware"
)

type Handler struct {
	handlerAuth      auth.HandlerAuth
	handlerAdmin     admin.HandlerAdmin
	handlerLesson    lesson.HandlerLesson
	adminUserHandler admin_user.AdminUserHandler
	middleware       middleware.Middleware
}

func NewHandler(handlerAuth auth.HandlerAuth, handlerAdmin admin.HandlerAdmin, handlerLesson lesson.HandlerLesson, adminUserHandler admin_user.AdminUserHandler, middleware middleware.Middleware) *Handler {
	return &Handler{
		handlerAuth:      handlerAuth,
		handlerAdmin:     handlerAdmin,
		handlerLesson:    handlerLesson,
		adminUserHandler: adminUserHandler,
		middleware:       middleware,
	}
}
