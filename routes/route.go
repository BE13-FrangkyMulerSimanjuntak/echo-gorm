package routes

import (
	"frangky/be13/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	e.GET("/users", controllers.GetAllUserController)
	e.POST("/users", controllers.AddUserController)
	e.GET("/users/:id", controllers.GetAllUserControllerId)
	e.PUT("/users/:id", controllers.UpddateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserContoller)
	return e
}
