package helper

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/keys"
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
	logger.Log.Debug(log_names.Helper, err.Error())
	code, msg := errs.HTTPError(err)

	c.JSON(code, dto.Result{
		Result: nil,
		Error:  new(msg),
	})
}

func GetAuthData(c *gin.Context) (sessionIdVal string, accessToken string, err error) {
	auth, _ := c.Get(keys.Authorization)
	sessionId, _ := c.Get(keys.XSessionID)

	sessionIdVal, err = GetHeaderVal(sessionId)
	if err != nil {
		return
	}

	accessToken, err = GetHeaderVal(auth)
	if err != nil {
		return
	}

	return
}
