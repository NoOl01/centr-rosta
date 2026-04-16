package personal_lesson

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/keys"
	"centr_rosta/internal/handler/dto"
	"centr_rosta/internal/handler/helper"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (aplh *AdminPersonalLessonHandler) GetPersonalLessonsRequests(c *gin.Context) {
	sessionIdVal, accessToken, err := helper.GetAuthData(c)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var resultPersonalLessons []dto.PersonalLesson
	personalLessons, err := aplh.upcl.GetPersonalLessonsRequests(ctx, sessionIdVal, accessToken)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	for _, personalLesson := range personalLessons {
		if *personalLesson.Status == "pending" {
			resultPersonalLessons = append(resultPersonalLessons, dto.PersonalLesson{
				ID:       personalLesson.ID,
				LessonID: personalLesson.LessonID,
				Lesson: dto.LessonData{
					ID:          personalLesson.Lesson.ID,
					Name:        personalLesson.Lesson.Name,
					Description: personalLesson.Lesson.Description,
				},
				UserID: *personalLesson.UserID,
				User: dto.User{
					ID:        personalLesson.User.ID,
					FirstName: &personalLesson.User.FirstName,
					LastName:  &personalLesson.User.LastName,
					Email:     &personalLesson.User.Email,
					Role:      personalLesson.User.Role,
				},
				TeacherID: personalLesson.TeacherID,
				Teacher: new(dto.User{
					ID:        personalLesson.Teacher.ID,
					FirstName: &personalLesson.Teacher.FirstName,
					LastName:  &personalLesson.Teacher.LastName,
					Email:     &personalLesson.Teacher.Email,
					Role:      personalLesson.Teacher.Role,
				}),
				EstimatedTimeFrom: *personalLesson.EstimatedTimeFrom,
				EstimatedTimeTo:   *personalLesson.EstimatedTimeTo,
				ExactTime:         personalLesson.ExactTime,
				Status:            *personalLesson.Status,
			})
		}
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: resultPersonalLessons,
		Error:  nil,
	})
}

func (aplh *AdminPersonalLessonHandler) ApprovePersonalLesson(c *gin.Context) {
	sessionIdVal, accessToken, err := helper.GetAuthData(c)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	var body dto.ApprovePersonalLesson
	if err := c.ShouldBindJSON(&body); err != nil {
		helper.HandleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	if err := aplh.upcl.ApprovePersonalLesson(ctx, sessionIdVal, accessToken, body.ID, body.Time); err != nil {
		helper.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: "ok",
		Error:  nil,
	})
}

func (aplh *AdminPersonalLessonHandler) CancelPersonalLesson(c *gin.Context) {
	sessionIdVal, accessToken, err := helper.GetAuthData(c)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	personalLessonIDStr := c.Query(keys.Id)
	if personalLessonIDStr == "" {
		helper.HandleError(c, errs.MissingQuery)
		return
	}

	personalLessonID, err := strconv.ParseInt(personalLessonIDStr, 10, 64)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	if err := aplh.upcl.CancelPersonalLesson(ctx, sessionIdVal, accessToken, personalLessonID); err != nil {
		helper.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: "ok",
		Error:  nil,
	})
}
