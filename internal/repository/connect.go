package repository

import (
	"centr_rosta/internal/config"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/internal/repository/models"
	"centr_rosta/pkg/logger"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		config.Env.DbHost,
		config.Env.DbUser,
		config.Env.DbPass,
		config.Env.DbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if err := db.AutoMigrate(&models.User{}, &models.Lesson{}, &models.GroupLessonSchedule{}, &models.GroupLessonSubscription{}, &models.PersonalLesson{}, &models.FavouriteLesson{}); err != nil {
		logger.Log.Error(log_names.Database, err.Error())
	}

	return db
}
