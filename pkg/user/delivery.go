package user

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/sarulabs/di"
)

type userHandler struct {
	usecase UseCase
}

// NewUserHTTPHandler group routes to user entity
func NewUserHTTPHandler(e *echo.Echo, ctn di.Container) {
	service := ctn.Get("user-usecase").(UseCase)

	handler := &userHandler{
		usecase: service,
	}

	g := e.Group("users")
	fmt.Println("[users] Echo Group: ", g)

	g.GET("", handler.GetAll)
}

func (u *userHandler) GetAll(c echo.Context) error {
	user := u.usecase.GetAll()

	println("userrrrrrr: ", user)

	return nil
}
