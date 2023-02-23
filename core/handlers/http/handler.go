package httphandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/matyv3/hexagonal-go-sst/core/domain"
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

func (c HTTPController) CreateTODO(ctx echo.Context) error {
	todo := new(domain.TODO)
	if err := ctx.Bind(todo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "There is an error with the body format").SetInternal(err)
	}

	result, err := c.service.CreateTODO(*todo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(err)
	}

	return ctx.JSON(201, result)
}

func (c HTTPController) GetTODOs(ctx echo.Context) error {
	todos, err := c.service.GetTODOs()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(err)
	}

	return ctx.JSON(200, todos)
}
