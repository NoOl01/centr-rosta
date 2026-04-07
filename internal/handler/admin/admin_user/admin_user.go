package admin_user

import (
	"centr_rosta/internal/consts/keys"
	"centr_rosta/internal/handler/dto"
	"centr_rosta/internal/handler/helper"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (hau *AdminUserHandler) GetUsers(c *gin.Context) {
	auth, _ := c.Get(keys.Authorization)
	sessionId, _ := c.Get(keys.XSessionID)

	sessionIdVal, err := helper.GetHeaderVal(sessionId)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	accessToken, err := helper.GetHeaderVal(auth)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	users, err := hau.uau.GetUsers(ctx, sessionIdVal, accessToken)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	var resUsers []dto.User

	for _, user := range users {
		resUsers = append(resUsers, dto.User{
			ID:        user.ID,
			FirstName: &user.FirstName,
			LastName:  &user.LastName,
			Email:     &user.Email,
			Role:      user.Role,
		})
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: resUsers,
		Error:  nil,
	})
}

func (hau *AdminUserHandler) UpdateUser(c *gin.Context) {
	auth, _ := c.Get(keys.Authorization)
	sessionId, _ := c.Get(keys.XSessionID)

	sessionIdVal, err := helper.GetHeaderVal(sessionId)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	accessToken, err := helper.GetHeaderVal(auth)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
}
