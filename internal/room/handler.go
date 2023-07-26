package room

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service RoomService
}

func NewHandler(svc RoomService) *Handler {
	return &Handler{service: svc}
}

func RegisterHandlers(e *echo.Echo, svc RoomService) {
	handler := NewHandler(svc)
	handler.registerRoomHandlers(e)
}

func (h *Handler) registerRoomHandlers(e *echo.Echo) {
	e.GET("/rooms", h.GetAllRoom)
	e.POST("/rooms/create", h.CreateRoom)
}

func (h *Handler) GetAllRoom(ctx echo.Context) error {
	//  validate
	//	r := ctx.Request()
	context := context.Background()

	rooms, err := h.service.GetAllRooms(context)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed get rooms")
	}

	return ctx.JSON(http.StatusOK, rooms)
}

func (h *Handler) CreateRoom(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "test")
}
