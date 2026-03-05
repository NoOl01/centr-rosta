package handler

import (
	"centr_rosta/internal/handler/middleware"

	"github.com/gin-gonic/gin"
)

func (h *handler) Router(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	{
		{
			auth := apiV1.Group("/auth")
			auth.POST("/register", h.handlerAuth.Register)
			auth.POST("/login", h.handlerAuth.Login)
			auth.POST("/refresh", h.handlerAuth.Refresh)
			auth.POST("/logout", h.handlerAuth.Logout)
			auth.GET("/check_access", middleware.AuthMiddleware(), h.handlerAuth.CheckAccess)
		}
		{
			admin := apiV1.Group("/admin")
			admin.GET("/stat", middleware.AuthMiddleware(), middleware.SessionMiddleware(), h.handlerAdmin.GetStatsByTimePeriod)
		}
	}
}
