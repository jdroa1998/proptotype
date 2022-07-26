package controller

import (
	"mvcPrototype/database"
	"mvcPrototype/model"
	response "mvcPrototype/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetStudents(context echo.Context) error {
	connection := database.GetConnection()
	students := new([]model.Student)

	connection.Preload("Country").Find(&students)
	var listStudentsResponse = []response.UserResponse{}
	for _, student := range *students {
		listStudentsResponse = append(listStudentsResponse, response.FromUserModelAndCountry(student, student.Country))
	}

	return context.JSON(http.StatusOK, listStudentsResponse)
}
