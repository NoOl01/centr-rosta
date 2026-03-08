package entity

type Payload struct {
	UserId string
	Role   string
}

type Login struct {
	Email    string
	Password string
}

type Refresh struct {
	RefreshToken string `json:"refresh_token"`
}
