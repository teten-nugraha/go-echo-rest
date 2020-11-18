package controllers

import (
	"github.com/labstack/echo"
	"github.com/teten-nugraha/simple-go-rest/models"
	"net/http"
)

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchAllPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
