package controller

import (
	service "mvcPrototype/service/student"

	"github.com/labstack/echo/v4"
)

type StudentController struct {
	echo.Context
}

func GetStudents(context echo.Context) error {
	return service.GetStudents(context)
}

// @Tags         Students
// @Title        Create Students
// @Description  Create Students
// @Param        body  body   request.CreateStudentRequest  true  "body for Students content"
// @Produce  json
// @Success 200 {object} response.UserResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /student [post]
func CreateStudent(context echo.Context) error {
	return service.CreateStudent(context)
}

func DeleteStudent(context echo.Context) error {
	return service.DeleteStudent(context)
}

func UpdateStudent(context echo.Context) error {
	return service.UpdateStudent(context)
}
