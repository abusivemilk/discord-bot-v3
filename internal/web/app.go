package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func App() *echo.Echo {
	e := echo.New()

	// API routes
	e.GET("/assignRoles/:id", assignRoles)

	return e
}

func assignRoles(c echo.Context) error {
	id := c.Param("id")
	SendToQueue(id)
	return c.JSON(http.StatusOK, nil)
}
