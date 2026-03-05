package middleware

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/keys"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/dto"
	"centr_rosta/pkg/logger"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader(keys.Authorization)
		if auth == "" {
			logger.Log.Debug(log_names.Middleware, errs.MissingHeader.Error())
			c.JSON(http.StatusUnauthorized, dto.Result{
				Result: nil,
				Error:  dto.Strconv(errs.MissingHeader.Error()),
			})
			c.Abort()
			return
		}
		if !strings.HasPrefix(auth, "Bearer ") {
			logger.Log.Debug(log_names.Middleware, errs.InvalidHeader.Error())
			c.JSON(http.StatusUnauthorized, dto.Result{
				Result: nil,
				Error:  dto.Strconv(errs.InvalidHeader.Error()),
			})
			c.Abort()
			return
		}
		auth = strings.TrimPrefix(auth, "Bearer ")

		c.Set(keys.Authorization, auth)
		c.Next()
	}
}

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId := c.GetHeader(keys.XSessionID)
		if sessionId == "" {
			logger.Log.Debug(log_names.Middleware, errs.MissingHeader.Error())
			c.JSON(http.StatusUnauthorized, dto.Result{
				Result: nil,
				Error:  dto.Strconv(errs.MissingHeader.Error()),
			})
			c.Abort()
			return
		}

		c.Set(keys.XSessionID, sessionId)
		c.Next()
	}
}
