package auth

import (
	"centr_rosta/internal/dto"
	"centr_rosta/internal/infra/session"
	"centr_rosta/internal/repository/auth"
)

type ServiceAuth interface {
	Register(user dto.User) error
	Login(user dto.Login) (string, string, error)
	Logout(token string) error
	Update(user dto.User) error
}

type serviceAuth struct {
	repo    auth.RepositoryAuth
	session session.RepositorySession
}

func (s serviceAuth) Register(user dto.User) error {
	//TODO implement me
	panic("implement me")
}

func (s serviceAuth) Login(user dto.Login) (string, string, error) {
	//TODO implement me
	panic("implement me")
}

func (s serviceAuth) Logout(token string) error {
	//TODO implement me
	panic("implement me")
}

func (s serviceAuth) Update(user dto.User) error {
	//TODO implement me
	panic("implement me")
}

func NewService(repo auth.RepositoryAuth, session session.RepositorySession) ServiceAuth {
	return &serviceAuth{
		repo:    repo,
		session: session,
	}
}
