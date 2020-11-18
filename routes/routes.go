package routes

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/simple-go-rest/controllers"
	"net/http"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "Hello this is from echo")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai)

	return e
}
