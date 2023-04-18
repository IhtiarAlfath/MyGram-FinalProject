package handler

import (
	"kominfo/h8-myGram-finalproject/dto"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

// Register godoc
// @Tags users
// @Description Create New user Data
// @ID register-user
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewUserRequest true "request body json"
// @Success 201 {object} dto.NewUserResponse
// @Router /users/register [post]
func (uh userHandler) Register(ctx *gin.Context) {
	var newUserRequest dto.NewUserRequest

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJSON := errs.NewUnproccessibleEntityError("invalid request body")

		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	response, err := uh.userService.CreateNewUser(newUserRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}


// Login godoc
// @Tags users
// @Description Login user Data
// @ID login-user
// @Accept json
// @Produce json
// @Param RequestBody body dto.LoginRequest true "request body json"
// @Success 201 {object} dto.LoginResponse
// @Router /users/login [post]
func (uh *userHandler) Login (ctx *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		errBindJSON := errs.NewUnproccessibleEntityError("invalid request body")

		ctx.JSON(errBindJSON.Status(), errBindJSON)
		return
	}

	response, err := uh.userService.Login(loginRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}