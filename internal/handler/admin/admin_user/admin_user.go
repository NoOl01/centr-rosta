package admin_user

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

func (hau *AdminUserHandler) GetUsers(c *gin.Context) {
	sessionIdVal, accessToken, err := helper.GetAuthData(c)
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

func (hau *AdminUserHandler) ResetPassword(c *gin.Context) {
	sessionIdVal, accessToken, err := helper.GetAuthData(c)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	userIDStr := c.Query(keys.Id)
	if userIDStr == "" {
		helper.HandleError(c, errs.MissingQuery)
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		helper.HandleError(c, errs.InvalidQuery)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	pass, err := hau.uau.ResetPassword(ctx, sessionIdVal, accessToken, userID)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: gin.H{
			"new_password": pass,
		},
		Error: nil,
	})
}

func (hau *AdminUserHandler) UpdateRole(c *gin.Context) {
	sessionIdVal, accessToken, err := helper.GetAuthData(c)
	if err != nil {
		helper.HandleError(c, err)
		return
	}

	var body dto.UpdateRole
	if err := c.ShouldBindJSON(&body); err != nil {
		helper.HandleError(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	if err := hau.uau.UpdateRole(ctx, sessionIdVal, accessToken, body.RoleName, body.UserID); err != nil {
		helper.HandleError(c, err)
	}

	c.JSON(http.StatusOK, dto.Result{
		Result: "ok",
		Error:  nil,
	})
}
