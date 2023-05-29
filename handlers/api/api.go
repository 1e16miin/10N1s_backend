package api

import (
	"github.com/10n1s-backend/handler/api/v1"
	"github.com/labstack/echo/v4"
)

func routes(e echo.Echo) {
	e.GET("v1/user/get", v1.getUsers)
}
