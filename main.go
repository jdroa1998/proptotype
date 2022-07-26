package main

import (
	"fmt"
	"mvcPrototype/common/utils"
	"mvcPrototype/controller"
	"mvcPrototype/database"
	_ "mvcPrototype/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	utils.LoggerInit()
	err := database.ConnectPostgres()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	router := echo.New()

	// Middleware
	router.Use(middleware.Recover())

	// Routes
	router.GET("/student", controller.GetStudents)
	router.POST("/student", controller.CreateStudent)
	router.DELETE("/student/:student", controller.DeleteStudent)
	router.PATCH("/student/:student", controller.UpdateStudent)
	router.GET("/health-check", controller.HealthCheck)
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	router.Logger.Fatal(router.Start(":1323"))
}
