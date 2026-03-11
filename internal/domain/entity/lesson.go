package entity

type LessonName struct {
	Name string
}

type Lesson struct {
	ID          *int64
	Name        string
	Description string
}
