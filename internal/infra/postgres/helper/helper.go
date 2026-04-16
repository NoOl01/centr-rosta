package helper

import (
	"centr_rosta/internal/domain/entity"
	"centr_rosta/internal/infra/postgres/models"
)

func UpdateUserStructBuild(dbUser models.User, updateUser entity.UpdateUser) (newUser models.User) {
	newUser = models.User{
		FirstName: isEmptyString(updateUser.FirstName, dbUser.FirstName),
		LastName:  isEmptyString(updateUser.LastName, dbUser.LastName),
		Email:     isEmptyString(updateUser.Email, dbUser.Email),
		Password:  isEmptyString(updateUser.Password, dbUser.Password),
		Role:      isEmptyString(updateUser.Role, dbUser.Role),
	}

	return
}

func UpdatePersonalLessonStructBuilder(dbPersonalLesson models.PersonalLesson, updatePersonalLesson entity.PersonalLesson) (newPersonalLesson models.PersonalLesson) {
	newPersonalLesson = models.PersonalLesson{
		ID:       isEmptyInt64(updatePersonalLesson.ID, dbPersonalLesson.ID),
		LessonID: isEmptyInt64(updatePersonalLesson.LessonID, dbPersonalLesson.LessonID),
		Lesson: models.Lesson{
			ID:          isEmptyInt64(updatePersonalLesson.Lesson.ID, dbPersonalLesson.Lesson.ID),
			Name:        updatePersonalLesson.Lesson.Name,
			Description: updatePersonalLesson.Lesson.Description,
			Schedules:   dbPersonalLesson.Lesson.Schedules,
		},
		UserID: *updatePersonalLesson.UserID,
		User: models.User{
			ID:        isEmptyInt64(updatePersonalLesson.User.ID, dbPersonalLesson.User.ID),
			FirstName: updatePersonalLesson.User.FirstName,
			LastName:  updatePersonalLesson.User.LastName,
			Email:     updatePersonalLesson.User.Email,
			Role:      isEmptyString(updatePersonalLesson.User.Role, dbPersonalLesson.User.Role),
			CreatedAt: dbPersonalLesson.CreatedAt,
		},
		TeacherID: new(isEmptyInt64(updatePersonalLesson.TeacherID, *dbPersonalLesson.TeacherID)),
		Teacher: new(models.User{
			ID:        isEmptyInt64(updatePersonalLesson.Teacher.ID, dbPersonalLesson.Teacher.ID),
			FirstName: updatePersonalLesson.Teacher.FirstName,
			LastName:  updatePersonalLesson.Teacher.LastName,
			Email:     updatePersonalLesson.Teacher.Email,
			Role:      isEmptyString(updatePersonalLesson.Teacher.Role, dbPersonalLesson.Teacher.Role),
			CreatedAt: dbPersonalLesson.CreatedAt,
		}),
		EstimatedTimeFrom: *updatePersonalLesson.EstimatedTimeFrom,
		EstimatedTimeTo:   *updatePersonalLesson.EstimatedTimeTo,
		ExactTime:         updatePersonalLesson.ExactTime,
		Status:            *updatePersonalLesson.Status,
		CreatedAt:         dbPersonalLesson.CreatedAt,
	}

	return
}

func isEmptyString(newVal *string, fallback string) string {
	if newVal == nil || *newVal == "" {
		return fallback
	}

	return *newVal
}

func isEmptyInt64(newVal *int64, fallback int64) int64 {
	if newVal == nil || *newVal == 0 {
		return fallback
	}

	return *newVal
}
