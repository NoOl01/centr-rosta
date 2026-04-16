package dto

import "time"

type Result struct {
	Result any     `json:"result,omitempty"`
	Error  *string `json:"error,omitempty"`
}

type PersonalLesson struct {
	ID                *int64     `json:"id,omitempty"`
	LessonID          *int64     `json:"lesson_id,omitempty"`
	Lesson            LessonData `json:"lesson,omitempty"`
	UserID            int64      `json:"user_id,omitempty"`
	User              User       `json:"user,omitempty"`
	TeacherID         *int64     `json:"teacher_id,omitempty"`
	Teacher           *User      `json:"teacher,omitempty"`
	EstimatedTimeFrom time.Time  `json:"estimated_time_from,omitempty"`
	EstimatedTimeTo   time.Time  `json:"estimated_time_to,omitempty"`
	ExactTime         *time.Time `json:"exact_time,omitempty"`
	Status            string     `json:"status,omitempty"`
}
