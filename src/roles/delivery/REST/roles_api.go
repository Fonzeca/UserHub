package REST

import (
	"context"
	"net/http"
	"time"

	"github.com/Fonzeka/Jame/src/domain"
	"github.com/Fonzeka/Jame/src/roles/usecase"
	"github.com/Fonzeka/Jame/src/security/jwt"
	"github.com/labstack/echo/v4"
)

type RolesApi struct {
	useCase usecase.RolesUseCase
}

//Constructor
func NewuserApi(useCase usecase.RolesUseCase) *RolesApi {
	return &RolesApi{useCase: useCase}
}

//Router
func (api *RolesApi) Router(e *echo.Echo) {
	e.POST("/admin/role", api.InsertRole, jwt.CheckInRole("admin"))
	e.DELETE("/admin/role", api.DeleteRole, jwt.CheckInRole("admin"))
	e.GET("/admin/roles", api.GetAllRoles, jwt.CheckInRole("admin"))
}

//Handlers ---------------

func (api *RolesApi) InsertRole(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var rol domain.Role
	c.Bind(&rol)

	err := api.useCase.InsertRole(ctx, rol)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (api *RolesApi) DeleteRole(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	val, err := c.FormParams()

	if err != nil {
		return err
	}

	name := val.Get("name")

	err = api.useCase.DeleteRole(ctx, name)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (api *RolesApi) GetAllRoles(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	data, err := api.useCase.GetAllRoles(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}
