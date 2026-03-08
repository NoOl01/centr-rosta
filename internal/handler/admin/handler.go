package admin

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/domain/usecase/admin"
	"centr_rosta/internal/handler/dto"
	"centr_rosta/pkg/logger"

	"github.com/gin-gonic/gin"
)

type HandlerAdmin struct {
	uad admin.IUseCaseAdmin
}

func NewHandlerAdmin(uad admin.IUseCaseAdmin) *HandlerAdmin {
	return &HandlerAdmin{
		uad: uad,
	}
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
