package admin

import (
	"centr_rosta/internal/domain/usecase/admin"
)

type HandlerAdmin struct {
	uad admin.IUseCaseAdmin
}

func NewHandlerAdmin(uad admin.IUseCaseAdmin) *HandlerAdmin {
	return &HandlerAdmin{
		uad: uad,
	}
}
