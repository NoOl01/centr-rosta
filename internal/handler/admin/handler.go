package admin

import (
	"centr_rosta/internal/usecase/admin"

	"github.com/gin-gonic/gin"
)

type HandlerAdmin interface {
	GetEmployees(c *gin.Context)
	GetLessons(c *gin.Context)
	GetSchedule(c *gin.Context)
	GetDefaultStats(c *gin.Context)
	GetStatsByTimePeriod(c *gin.Context)
}

type handlerAdmin struct {
	uad admin.UseCaseAdmin
}

func (ha *handlerAdmin) GetEmployees(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (ha *handlerAdmin) GetLessons(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (ha *handlerAdmin) GetSchedule(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (ha *handlerAdmin) GetDefaultStats(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewHandlerAdmin(uad admin.UseCaseAdmin) HandlerAdmin {
	return &handlerAdmin{
		uad: uad,
	}
}
