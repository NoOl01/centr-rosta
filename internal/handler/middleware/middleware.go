package middleware

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/keys"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/dto"
	"centr_rosta/pkg/logger"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AuthMiddleware() gin.HandlerFunc
	SessionMiddleware() gin.HandlerFunc
}

type middleware struct{}

func NewMiddleware() Middleware {
	return &middleware{}
}

func (m *middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader(keys.Authorization)
		if auth == "" {
			code, msg := errs.HTTPError(errs.MissingHeader)

			logger.Log.Debug(log_names.Middleware, msg)
			c.JSON(code, dto.Result{
				Result: nil,
				Error:  dto.Strconv(msg),
			})
			c.Abort()
			return
		}
		if !strings.HasPrefix(auth, "Bearer ") {
			code, msg := errs.HTTPError(errs.InvalidHeader)

			logger.Log.Debug(log_names.Middleware, msg)
			c.JSON(code, dto.Result{
				Result: nil,
				Error:  dto.Strconv(msg),
			})
			c.Abort()
			return
		}
		auth = strings.TrimPrefix(auth, "Bearer ")

		c.Set(keys.Authorization, auth)
		c.Next()
	}
}

func (m *middleware) SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId := c.GetHeader(keys.XSessionID)
		if sessionId == "" {
			code, msg := errs.HTTPError(errs.MissingHeader)

			logger.Log.Debug(log_names.Middleware, msg)
			c.JSON(code, dto.Result{
				Result: nil,
				Error:  dto.Strconv(msg),
			})
			c.Abort()
			return
		}

		c.Set(keys.XSessionID, sessionId)
		c.Next()
	}
}
