package repository

import (
	"centr_rosta/internal/config"
	"centr_rosta/internal/consts"
	"centr_rosta/internal/repository/models"
	"centr_rosta/pkg/logger"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4",
		config.Env.DbUser, config.Env.DbPass, config.Env.DbPort, config.Env.DbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Error(consts.Database, err.Error())
		return nil
	}

	if err := db.AutoMigrate(&models.User{}, &models.Lesson{}, &models.GroupLessonSchedule{}, &models.GroupLessonSubscription{}, &models.PersonalLesson{}, &models.FavouriteLesson{}); err != nil {
		logger.Log.Error(consts.Database, err.Error())
	}

	return db
}
