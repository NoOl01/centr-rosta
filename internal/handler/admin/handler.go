package admin

import (
	"centr_rosta/internal/domain/usecase/admin"
)

type HandlerAdmin struct {
	uad admin.UseCaseAdmin
}

func NewHandlerAdmin(uad admin.UseCaseAdmin) *HandlerAdmin {
	return &HandlerAdmin{
		uad: uad,
	}
}
