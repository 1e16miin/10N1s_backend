package route

import (
	"context"
	"errors"
	"fmt"

	"github.com/10n1s-backend/cmd/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"net/http"
)

const (
	BasicAuthHeader = "Authorization"
	AccountHeader   = "Account"
	WorkflowHeader  = "Id"
	EndpointHeader  = "Endpoint"
)

var (
	ErrWrongRequestBasicAuth  = errors.New("unable to parse request's basic auth")
	ErrWrongRequestAccountID  = errors.New("unable to parse request's AccountID")
	ErrWrongRequestWorkflowID = errors.New("unable to parse request's WorkflowID")
	WrongRequestCode          = http.StatusBadRequest

	ErrInternal       = errors.New("internal server error")
	InternalErrorCode = http.StatusInternalServerError

	StatusServiceUnavailableCode = http.StatusServiceUnavailable

	ErrNoAccountHeader  = errors.New("http request has no account header")
	ErrNoWorkflowHeader = errors.New("http request has no workflow header")
)

type (
	echoRouter struct {
		route      *echo.Echo
		config     EchoConfig
		controller controller.Interface
	}

	EchoConfig struct {
		Port string `config:"port"`
	}
)

func NewEchoRouter(ctx context.Context, config EchoConfig, controller controller.Interface) (*echoRouter, error) {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	echoRouter := &echoRouter{route: e, config: config, controller: controller}

	err := echoRouter.AddHandler(ctx)
	if err != nil {
		return nil, err
	}

	return echoRouter, nil

}

func (e *echoRouter) AddHandler(ctx context.Context) error {
	e.route.Use(middleware.Logger())
	e.route.Use(middleware.Recover())
	e.route.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
	}))

	e.route.GET("test", e.testGetHandler)

	return nil
}

func (e *echoRouter) Start() error {
	return e.route.Start(fmt.Sprintf(":%s", "8080"))
}

func (e *echoRouter) testGetHandler(c echo.Context) error {
	//ctx := c.Request().Context()

	return c.String(http.StatusOK, "")
}
