package controller

import (
	"fmt"
	"mvcPrototype/database"
	"mvcPrototype/model"
	request "mvcPrototype/request"
	response "mvcPrototype/response"
	email "mvcPrototype/service/email"
	pensum "mvcPrototype/service/pensum"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateStudent(context echo.Context) error {
	if pensum.Exists() {
		requestStudent := new(request.CreateStudentRequest)
		if err := context.Bind(&requestStudent); err != nil {
			return err
		}

		student := request.CreateStudentRequestToStudent(*requestStudent)
		fmt.Println(requestStudent)
		connection := database.GetConnection()
		connection.Save(&student)
		country := new(model.Country)
		connection.First(&country, student.CountryId)
		context.JSON(http.StatusCreated, response.FromUserModelAndCountry(student, *country))

		if err := email.Send(); err != nil {
			panic(err)
		}
		return nil
	} else {
		return nil
	}
}
