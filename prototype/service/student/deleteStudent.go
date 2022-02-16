package controller

import (
	"mvcPrototype/database"
	"mvcPrototype/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteStudent(context echo.Context) error {
	studentId := context.Param("student")
	if studentId == "" {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"code":    "error.InvalidInput",
			"message": "Parameter 'studentId' is empty or invalid.",
		})
	}

	connection := database.GetConnection()
	connection.Unscoped().Delete(&model.Student{}, studentId)

	return nil
}
