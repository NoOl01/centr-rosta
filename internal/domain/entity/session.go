package entity

type Session struct {
	UserID       string
	DeviceToken  string
	AccessToken  string
	RefreshToken string
}
