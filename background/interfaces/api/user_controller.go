package controllers

import (
	domain "github.com/YuminosukeSato/slopegram/background/domain"
	"github.com/YuminosukeSato/slopegram/background/interfaces/database"
	usecase "github.com/YuminosukeSato/slopegram/background/usecase"

	"github.com/labstack/echo"
)

type UserController struct {
    Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
    return &UserController{
        Interactor: usecase.UserInteractor{
            UserRepository: &database.UserRepository{
                SqlHandler: sqlHandler,
            },
        },
    }
}

func (controller *UserController) Create(c echo.Context) {
    u := domain.User{}
    c.Bind(&u)
    controller.Interactor.Add(u)
    createdUsers := controller.Interactor.GetInfo()
    c.JSON(201, createdUsers)
    return
}

func (controller *UserController) GetUser() []domain.User {
    res := controller.Interactor.GetInfo()
    return res
}

func (controller *UserController) Delete(id string) {
    controller.Interactor.Delete(id)
}
