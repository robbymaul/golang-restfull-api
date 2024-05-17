package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	service "github.com/robbymaul/golang-swagger.git/service/user"
	web "github.com/robbymaul/golang-swagger.git/web"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Register(c *gin.Context) {
	request := new(web.UserRegisterRequest)
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	response, _ := controller.UserService.Register(request)

	c.JSON(http.StatusOK, web.ResponseStatusOk{
		Status: true,
		Data:   response,
	})

}
