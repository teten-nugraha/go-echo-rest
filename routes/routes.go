package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/teten-nugraha/simple-go-rest/controllers"
	"github.com/teten-nugraha/simple-go-rest/middlewares"
	"net/http"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "Hello this is from echo")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai, middlewares.IsAuthenticated)
	e.POST("/pegawai", controllers.StorePegawai, middlewares.IsAuthenticated)
	e.PUT("/pegawai", controllers.UpdatePegawai, middlewares.IsAuthenticated)
	e.DELETE("/pegawai", controllers.DeletePegawai, middlewares.IsAuthenticated)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.GET("/login", controllers.CheckLogin)

	return e
}
