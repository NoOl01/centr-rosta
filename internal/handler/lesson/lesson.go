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

	c.JSON(http.StatusOK, dto.Result{
		Error:  nil,
		Result: lessons,
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

	c.JSON(http.StatusOK, dto.Result{
		Error:  nil,
		Result: lesson,
	})
}
