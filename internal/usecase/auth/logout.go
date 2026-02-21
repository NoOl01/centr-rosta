package auth

import "context"

func (ua *useCaseAuth) Logout(ctx context.Context, sessionID string) error {
	return ua.session.Delete(ctx, sessionID)
}
