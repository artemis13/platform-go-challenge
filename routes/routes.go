package routes

import (
	"github.com/artemis13/platform-go-challenge/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/users/:id/favorites", handlers.GetUserFavorites)
	e.POST("/users/:id/favorites", handlers.AddUserFavorite)
	e.DELETE("/users/:id/favorites/:asset_id", handlers.RemoveUserFavorite)
	e.PUT("/users/:id/favorites/:asset_id", handlers.EditUserFavorite)
}
