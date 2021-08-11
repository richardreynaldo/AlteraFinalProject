package users

import (
	"finalProject/business/users"
	"finalProject/controllers"
	"finalProject/controllers/users/request"
	"net/http"

	"github.com/labstack/echo"
)

type UserController struct {
	UserUseCase users.UseCase
}

func NewUserController(e *echo.Echo, userUC users.UseCase) {
	controller := &UserController{
		UserUseCase: userUC,
	}

	news := e.Group("user")
	news.POST("/login", controller.Login)
}

func (controller *UserController) Login(c echo.Context) error {
	var userLogin request.LoginUser
	if err := c.Bind(&userLogin); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	user, err := controller.UserUseCase.Login(c.Request().Context(), userLogin.Email, userLogin.Password)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, user)
}
