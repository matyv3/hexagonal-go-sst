package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	httphandler "github.com/matyv3/hexagonal-go-sst/core/handlers/http"
	"github.com/matyv3/hexagonal-go-sst/core/repositories"
	"github.com/matyv3/hexagonal-go-sst/core/services"
)

func StartHTTPServer() {
	fmt.Println("starting http server...")

	repository := repositories.CreateTODORepository()
	service := services.CreateTODOService(repository)
	controller := httphandler.CreateHTTPController(service)

	server := echo.New()
	// Routes
	server.GET("/todos", controller.GetTODOs)
	server.POST("/todos", controller.CreateTODO)

	// Start server
	server.Logger.Fatal(server.Start(":4000"))
}
