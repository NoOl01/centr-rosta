package entity

type Transaction struct {
	UserID   int64
	User     User
	Amount   float64
	Type     string
	LessonID int64
	Lesson   LessonName
}
