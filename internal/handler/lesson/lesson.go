package lesson

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/keys"
	"centr_rosta/internal/domain/entity"
	"centr_rosta/internal/handler/dto"
	"centr_rosta/internal/handler/helper"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (hl *HandlerLesson) CreateLesson(c *gin.Context) {
	sessionIdVal, accessToken, err := helper.GetAuthData(c)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	var body dto.LessonData
	if err := c.ShouldBindJSON(&body); err != nil {
		helper.HandleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	lesson := entity.Lesson{
		Name:        body.Name,
		Description: body.Description,
	}

	if err := hl.ul.CreateLesson(ctx, sessionIdVal, accessToken, &lesson); err != nil {
		helper.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: "ok",
		Error:  nil,
	})
}

func (hl *HandlerLesson) UpdateLesson(c *gin.Context) {
	sessionIdVal, accessToken, err := helper.GetAuthData(c)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	var body dto.LessonData
	if err := c.ShouldBindJSON(&body); err != nil {
		helper.HandleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	lesson := entity.Lesson{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
	}

	if err := hl.ul.UpdateLesson(ctx, sessionIdVal, accessToken, &lesson); err != nil {
		helper.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: "ok",
		Error:  nil,
	})
}

func (hl *HandlerLesson) GetLessons(c *gin.Context) {
	lessons, err := hl.ul.GetLessons()
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	resLessons := make([]dto.LessonData, 0, len(lessons))

	for _, l := range lessons {
		resLessons = append(resLessons, dto.LessonData{
			ID:          l.ID,
			Name:        l.Name,
			Description: l.Description,
		})
	}

	c.JSON(http.StatusOK, dto.Result{
		Error:  nil,
		Result: resLessons,
	})
}

func (hl *HandlerLesson) GetLessonByID(c *gin.Context) {
	lessonID, err := strconv.ParseInt(c.Param(keys.Id), 10, 64)
	if err != nil {
		helper.HandleError(c, errs.InternalError)
		return
	}

	lesson, err := hl.ul.GetLessonByID(lessonID)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	resLesson := dto.LessonData{
		ID:          lesson.ID,
		Name:        lesson.Name,
		Description: lesson.Description,
	}

	c.JSON(http.StatusOK, dto.Result{
		Error:  nil,
		Result: resLesson,
	})
}
