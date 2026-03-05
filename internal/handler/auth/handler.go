package auth

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/dto"
	"centr_rosta/internal/usecase/auth"
	"centr_rosta/pkg/logger"

	"github.com/gin-gonic/gin"
)

type HandlerAuth interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Refresh(c *gin.Context)
	Logout(c *gin.Context)
	CheckAccess(c *gin.Context)
}

type handlerAuth struct {
	ua auth.UseCaseAuth
}

func NewHandlerAuth(ua auth.UseCaseAuth) HandlerAuth {
	return &handlerAuth{ua: ua}
}

func getHeaderVal(headerValue any) (string, error) {
	value, ok := headerValue.(string)
	if !ok {
		return "", errs.MissingHeader
	}

	return value, nil
}

func handleError(c *gin.Context, err error) {
	logger.Log.Debug(log_names.AuthHandler, err.Error())
	code, msg := errs.HTTPError(err)
	c.JSON(code, dto.Result{
		Result: nil,
		Error:  dto.Strconv(msg),
	})
}
