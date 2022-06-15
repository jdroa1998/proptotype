package controller

import (
	"fmt"
	"mvcPrototype/database"
	"mvcPrototype/model"
	request "mvcPrototype/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateStudent(context echo.Context) error {
	studentId := context.Param("student")
	if studentId == "" {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"code":    "error.InvalidInput",
			"message": "Parameter 'studentId' is empty or invalid.",
		})
	}

	requestStudent := new(request.UpdateStudentRequest)
	if err := context.Bind(&requestStudent); err != nil {
		return err
	}

	connection := database.GetConnection()
	student := new(model.Student)
	connection.Find(&student, studentId)

	request.UpdateStudentRequestToStudent(*requestStudent, student)

	update := connection.Save(student)
	if update.Error != nil {
		fmt.Println(update.Error)
		return context.JSON(http.StatusBadRequest, map[string]string{
			"code":    "error.updateError",
			"message": "Something went wrong.",
		})
	}

	return nil
}
