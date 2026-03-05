package lesson

import "gorm.io/gorm"

type RepositoryLesson interface {
}

type repositoryLesson struct {
	Db *gorm.DB
}

func NewRepositoryLesson(db *gorm.DB) RepositoryLesson {
	return &repositoryLesson{Db: db}
}
