package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRooms(ctx echo.Context) error {
	//validate
	rooms, err := h.service.GetRooms()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed get rooms")
	}

	return ctx.JSON(http.StatusOK, rooms)
}

func (h *Handler) CreateRoom(ctx echo.Context) error {
	//validate
	//r := ctx.Request()

	//rooms, err := h.service.CreateRoom()
	// if err != nil {
	// 	return ctx.String(http.StatusInternalServerError, "Failed get rooms")
	// }

	return ctx.JSON(http.StatusOK, "test")
}

func (h *Handler) registerRoomHandlers(e *echo.Echo) {
	e.GET("/rooms", h.GetRooms)
	e.POST("/rooms/create", h.CreateRoom)
}
