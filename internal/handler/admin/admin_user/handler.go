package admin_user

import "centr_rosta/internal/domain/usecase/admin/admin_user"

type AdminUserHandler struct {
	uau admin_user.UseCaseAdminUser
}

func NewAdminUserHandler(uau admin_user.UseCaseAdminUser) *AdminUserHandler {
	return &AdminUserHandler{
		uau: uau,
	}
}
