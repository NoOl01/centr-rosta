package entity

type User struct {
	ID        *int64
	FirstName string
	LastName  string
	Email     string
	Password  *string
	Role      *string
}

type UpdateUser struct {
	FirstName *string
	LastName  *string
	Email     *string
	Password  *string
	Role      *string
}
