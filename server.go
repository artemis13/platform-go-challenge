package main

import (
	"net/http"

	"github.com/artemis13/platform-go-challenge/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "GWI Go Platform Challenge running!\n")
	})

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
