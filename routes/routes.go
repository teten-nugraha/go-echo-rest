package routes

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/simple-go-rest/controllers"
	"github.com/teten-nugraha/simple-go-rest/middlewareConf"
	"net/http"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "Hello this is from echo")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai, middlewareConf.IsAuthenticated)
	e.POST("/pegawai", controllers.StorePegawai, middlewareConf.IsAuthenticated)
	e.PUT("/pegawai", controllers.UpdatePegawai, middlewareConf.IsAuthenticated)
	e.DELETE("/pegawai", controllers.DeletePegawai, middlewareConf.IsAuthenticated)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.GET("/login", controllers.CheckLogin)

	return e
}
