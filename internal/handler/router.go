package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Router(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	{
		{
			auth := apiV1.Group("/auth")
			auth.POST("/register", h.handlerAuth.Register)
			auth.POST("/login", h.handlerAuth.Login)
			auth.POST("/refresh", h.middleware.SessionMiddleware(), h.handlerAuth.Refresh)
			auth.POST("/logout", h.middleware.SessionMiddleware(), h.handlerAuth.Logout)
			auth.GET("/check_access", h.middleware.AuthMiddleware(), h.middleware.SessionMiddleware(), h.handlerAuth.CheckAccess)
		}
		{
			admin := apiV1.Group("/admin")
			admin.GET("/stat", h.middleware.AuthMiddleware(), h.middleware.SessionMiddleware(), h.handlerAdmin.GetStatsByTimePeriod)
			{
				user := apiV1.Group("/user")
				user.GET("/")
				user.POST("/")
				user.PATCH("/")
			}
		}
		{
			lesson := apiV1.Group("/lesson")
			lesson.GET("/")
			lesson.GET("/favourite")
			lesson.GET("/group")
			lesson.GET("/personal")
			lesson.POST("/subscribe")
			lesson.DELETE("/cancel")
		}
		{
			schedule := apiV1.Group("/schedule")
			schedule.GET("/")
		}
	}
}
