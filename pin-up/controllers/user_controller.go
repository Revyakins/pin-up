package controllers

import (
	"github.com/gin-gonic/gin"
	"go-app/pin-up/domain/user"
	"go-app/pin-up/services/user_service"
	"net/http"
)

type UserControllerInterface interface {
	CreateUser(ctx *gin.Context)
}

type UserInput struct {
	Id  string `json:"id"`
	Balance string   `json:"balance"`
}

type UserOutPut struct {
	error string
}

type userController struct {
	us user_service.UserServiceInterface
}

func (c *userController) CreateUser(ctx *gin.Context)  {

	var userInput UserInput
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		HTTPRes(ctx, http.StatusBadRequest, err.Error())
		return
	}

	u := user.User{
		Id: userInput.Id,
		Balance: userInput.Balance,
	}

	if _, err := c.us.CreateUser(&u); err != nil {
		HTTPRes(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	HTTPRes(ctx, http.StatusOK, "")
}

func NewUserController(us user_service.UserServiceInterface) UserControllerInterface {
	return &userController{us: us}
}
