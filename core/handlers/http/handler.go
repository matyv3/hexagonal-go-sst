package httphandler

import (
	"github.com/labstack/echo/v4"
	"github.com/matyv3/hexagonal-go-sst/core/ports"
)

type HTTPController struct {
	service ports.TODOService
}

func CreateHTTPController(service ports.TODOService) HTTPController {
	return HTTPController{
		service: service,
	}
}

func (c HTTPController) CreateTODO(ctx echo.HandlerFunc) {
}

func (c HTTPController) GetTODOS(ctx echo.HandlerFunc) {
}
