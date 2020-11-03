package routes

import (
	"github.com/labstack/echo"
	"net/http"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "Hello this is from echo")
	})

	return e
}
