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
			admin := apiV1.Group("/admin", h.middleware.AuthMiddleware(), h.middleware.SessionMiddleware())
			admin.GET("/stat", h.handlerAdmin.GetStatsByTimePeriod)
			{
				{
					user := admin.Group("/user")
					user.GET("/", h.adminUserHandler.GetUsers)
					user.POST("/reset-pass", h.adminUserHandler.ResetPassword)
					user.PATCH("/", h.adminUserHandler.UpdateRole)
				}
				{
					lesson := admin.Group("/lesson")
					lesson.GET("/", h.handlerLesson.GetLesson)
					lesson.POST("/", h.handlerLesson.CreateLesson)
					lesson.PATCH("/", h.handlerLesson.UpdateLesson)
				}
				{
					personalLesson := admin.Group("/personal-lesson")
					personalLesson.GET("/")
					personalLesson.POST("/approve")
					personalLesson.DELETE("/cancel")
				}
			}
		}
		{
			lesson := apiV1.Group("/lesson")
			lesson.GET("/", h.handlerLesson.GetLesson)
			lesson.GET("/favourite", h.middleware.AuthMiddleware(), h.middleware.SessionMiddleware())
			lesson.POST("/favourite", h.middleware.AuthMiddleware(), h.middleware.SessionMiddleware())
		}
		{
			schedule := apiV1.Group("/schedule", h.middleware.AuthMiddleware(), h.middleware.SessionMiddleware())
			schedule.GET("/")
			schedule.POST("/group")
			schedule.POST("/personal")
			schedule.DELETE("/cancel")
		}
		{
			payment := apiV1.Group("/payment", h.middleware.AuthMiddleware(), h.middleware.SessionMiddleware())
			payment.POST("/create")
		}
	}
}
