package handler

import "github.com/gin-gonic/gin"

func (h *handler) Router(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	{
		auth := apiV1.Group("/auth")
		{
			auth.POST("/register", h.handlerAuth.Register)
			auth.POST("/login", h.handlerAuth.Login)
			auth.POST("/refresh", h.handlerAuth.Refresh)
			auth.POST("/logout", h.handlerAuth.Logout)
		}
	}
}
