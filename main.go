package main

import (
	"fmt"
	"net/http"
	"requestaggregator/app"
	aggregatorController "requestaggregator/app/aggregator"
	"requestaggregator/config"
	"requestaggregator/internal/command"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Request Aggregator Project")
	})

	e.POST("/v1/query", aggregatorController.HandleGetAggregatorData)
	e.POST("/v1/insert", aggregatorController.HandleInsertAggregatorData)

	// Server
	app.SetCommands()
	command.RunApp()
	appConfig := config.GetConfig()
	fmt.Println("Starting server at ", appConfig.Port)
	e.Run(standard.New(fmt.Sprintf(":%d", appConfig.Port)))

}
