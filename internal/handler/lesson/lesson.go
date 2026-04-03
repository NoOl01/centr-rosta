package lesson

import (
	"centr_rosta/internal/consts/errs"
	"centr_rosta/internal/consts/keys"
	"centr_rosta/internal/handler/dto"
	"centr_rosta/internal/handler/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hl *HandlerLesson) GetLessons(c *gin.Context) {
	lessons, err := hl.ul.GetLessons()
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	resLessons := make([]dto.LessonData, 0, len(lessons))

	for _, l := range lessons {
		resLessons = append(resLessons, dto.LessonData{
			ID:          *l.ID,
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
		ID:          *lesson.ID,
		Name:        lesson.Name,
		Description: lesson.Description,
	}

	c.JSON(http.StatusOK, dto.Result{
		Error:  nil,
		Result: resLesson,
	})
}
