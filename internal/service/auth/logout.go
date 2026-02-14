package auth

import "context"

func (s *serviceAuth) Logout(ctx context.Context, sessionID string) error {
	return s.session.Delete(ctx, sessionID)
}
