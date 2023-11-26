package controllers

import (
	"github.com/labstack/echo/v4"
)

func (s *service) RegisterHealth(g *echo.Group) {
	g.GET("/", s.health_getHealth)
	g.GET("/database", s.health_databaseHealth)

}

func (s *service) health_getHealth(c echo.Context) error {
	return c.String(200, "OK")
}

func (s *service) health_databaseHealth(c echo.Context) error {
	return s.db.Health()
}
