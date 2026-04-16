package dto

type User struct {
	ID        *int64  `json:"ID,omitempty"`
	FirstName *string `json:"first_name,omitempty" json:"firstName,omitempty"`
	LastName  *string `json:"last_name,omitempty" json:"lastName,omitempty"`
	Email     *string `json:"email,omitempty" json:"email,omitempty"`
	Password  *string `json:"password,omitempty" json:"password,omitempty"`
	Role      *string `json:"role,omitempty" json:"role,omitempty"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Refresh struct {
	RefreshToken string `json:"refresh_token"`
}

type UpdateRole struct {
	UserID   int64  `json:"user_id"`
	RoleName string `json:"role_name"`
}

type TransactionStats struct {
	TotalAmount float64       `json:"total_amount"`
	Transaction []Transaction `json:"transaction"`
}
type Transaction struct {
	UserID   int64    `json:"user_id"`
	User     UserInfo `json:"user"`
	Amount   float64  `json:"amount"`
	Type     string   `json:"type"`
	LessonID int64    `json:"lesson_id"`
	Lesson   Lesson   `json:"lesson"`
}

type UserInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Lesson struct {
	Name string `json:"name"`
}

type LessonData struct {
	ID          *int64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ApprovePersonalLesson struct {
	ID   int64  `json:"id"`
	Time string `json:"time"`
}
