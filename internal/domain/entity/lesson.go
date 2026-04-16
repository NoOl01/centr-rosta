package entity

import "time"

type LessonName struct {
	Name string
}

type Lesson struct {
	ID          *int64
	Name        string
	Description string
}

type PersonalLesson struct {
	ID                *int64
	LessonID          *int64
	Lesson            *Lesson
	UserID            *int64
	User              *User
	TeacherID         *int64
	Teacher           *User
	EstimatedTimeFrom *time.Time
	EstimatedTimeTo   *time.Time
	ExactTime         *time.Time
	Status            *string
}
