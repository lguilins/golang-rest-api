package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lguilins/oraculosweb-api/infra"
	"github.com/sarulabs/di"

	userHTTP "github.com/lguilins/oraculosweb-api/pkg/user"
)

// RunHTTPServer runs a Rest HTTP Server
func RunHTTPServer(ctn di.Container) error {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Application running...")
	})

	userHTTP.NewUserHTTPHandler(e, ctn)

	e.Start(fmt.Sprintf(":%v", infra.AppPort))

	return nil
}
