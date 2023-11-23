package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterHealth(g *echo.Group) {
	g.GET("/", getHealth)
}

func getHealth(c echo.Context) error {
	return c.String(200, "OK")
}
