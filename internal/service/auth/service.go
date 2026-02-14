package auth

import (
	"centr_rosta/internal/dto"
	"centr_rosta/internal/infra/session"
	"centr_rosta/internal/repository/auth"
	"context"
)

type ServiceAuth interface {
	Register(ctx context.Context, user dto.User) (string, string, string, error)
	Login(ctx context.Context, user dto.Login) (string, string, string, error)
	Refresh(ctx context.Context, sessionID string, refreshData dto.Refresh) (string, string, error)
	Logout(token string) error
	Update(user dto.User) error
}

type serviceAuth struct {
	repo    auth.RepositoryAuth
	session session.RepositorySession
}

func (s *serviceAuth) Logout(token string) error {
	//TODO implement me
	panic("implement me")
}

func (s *serviceAuth) Update(user dto.User) error {
	//TODO implement me
	panic("implement me")
}

func NewService(repo auth.RepositoryAuth, session session.RepositorySession) ServiceAuth {
	return &serviceAuth{
		repo:    repo,
		session: session,
	}
}
