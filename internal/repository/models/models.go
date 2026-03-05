package models

import "time"

type User struct {
	ID                   int64                     `gorm:"primaryKey;autoIncrement"`
	FirstName            string                    `gorm:"not null"`
	LastName             string                    `gorm:"not null"`
	Email                string                    `gorm:"not null;uniqueIndex"`
	Password             string                    `gorm:"not null"`
	Role                 string                    `gorm:"not null;default:'user'"`
	CreatedAt            time.Time                 `gorm:"not null"`
	FavouriteLessons     []FavouriteLesson         `gorm:"foreignKey:UserID"`
	Subscriptions        []GroupLessonSubscription `gorm:"foreignKey:UserID"`
	TeachingGroupLessons []GroupLessonSchedule     `gorm:"foreignKey:TeacherID"`
	PersonalLessons      []PersonalLesson          `gorm:"foreignKey:UserID"`
}

type Lesson struct {
	ID          int64                 `gorm:"primaryKey;autoIncrement"`
	Name        string                `gorm:"not null"`
	Description string                `gorm:"not null"`
	Schedules   []GroupLessonSchedule `gorm:"foreignKey:LessonID"`
}

type GroupLessonSchedule struct {
	ID            int64                     `gorm:"primaryKey;autoIncrement"`
	LessonID      int64                     `gorm:"not null"`
	Lesson        Lesson                    `gorm:"foreignKey:LessonID"`
	TeacherID     int64                     `gorm:"not null"`
	Teacher       User                      `gorm:"foreignKey:TeacherID"`
	Time          time.Time                 `gorm:"not null"`
	Subscriptions []GroupLessonSubscription `gorm:"foreignKey:GroupLessonScheduleID"`
}

type GroupLessonSubscription struct {
	ID                    int64               `gorm:"primaryKey;autoIncrement"`
	UserID                int64               `gorm:"not null"`
	User                  User                `gorm:"foreignKey:UserID"`
	GroupLessonScheduleID int64               `gorm:"not null"`
	GroupLessonSchedule   GroupLessonSchedule `gorm:"foreignKey:GroupLessonScheduleID"`
	CreatedAt             time.Time           `gorm:"autoCreateTime"`
}

type PersonalLesson struct {
	ID            int64  `gorm:"primaryKey;autoIncrement"`
	LessonID      int64  `gorm:"not null"`
	Lesson        Lesson `gorm:"foreignKey:LessonID"`
	UserID        int64  `gorm:"not null"`
	User          User   `gorm:"foreignKey:UserID"`
	TeacherID     *int64
	Teacher       *User     `gorm:"foreignKey:TeacherID"`
	EstimatedTime time.Time `gorm:"not null"`
	ExactTime     *time.Time
	Status        string    `gorm:"not null;default:'pending'"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

type FavouriteLesson struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	LessonID int64  `gorm:"not null"`
	Lesson   Lesson `gorm:"foreignKey:LessonID"`
	UserID   int64  `gorm:"not null"`
	User     User   `gorm:"foreignKey:UserID"`
}

type Transaction struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	UserID    int64     `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	Amount    float64   `gorm:"not null"`
	Type      string    `gorm:"not null"`
	LessonID  int64     `gorm:"not null"`
	Lesson    Lesson    `gorm:"foreignKey:LessonID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
