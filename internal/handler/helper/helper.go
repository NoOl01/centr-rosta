package helper

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/handler/dto"
	"centr_rosta/pkg/logger"

	"github.com/gin-gonic/gin"
)

func GetHeaderVal(headerValue any) (string, error) {
	value, ok := headerValue.(string)
	if !ok {
		return "", errs.MissingHeader
	}

	return value, nil
}

func HandleError(c *gin.Context, err error) {
	logger.Log.Debug(log_names.AuthHandler, err.Error())
	code, msg := errs.HTTPError(err)
	c.JSON(code, dto.Result{
		Result: nil,
		Error:  dto.Strconv(msg),
	})
}
