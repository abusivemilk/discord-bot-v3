package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func App() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://vatusa.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// API routes
	e.GET("/assignRoles/:id", assignRoles)
	e.POST("/assignRoles/:id", assignRoles)

	return e
}

func assignRoles(c echo.Context) error {
	id := c.Param("id")
	SendToQueue(id)
	return c.JSON(http.StatusOK, nil)
}
