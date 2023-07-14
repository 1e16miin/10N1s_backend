package handler

import (
	"github.com/10n1s-backend/internal/service"
	"github.com/10n1s-backend/pkg/logger"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
	logger  logger.Logger
}

func NewHandler(svc *service.Service, logger logger.Logger) *Handler {
	return &Handler{service: svc, logger: logger}
}

func RegisterHandlers(e *echo.Echo, svc *service.Service, logger logger.Logger) {
	handler := NewHandler(svc, logger)

	handler.registerRoomHandlers(e)
	//handler.registerUserHandlers(e)
}
