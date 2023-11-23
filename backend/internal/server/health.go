package server

import (
	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterHealth(g *echo.Group) {
	g.GET("/", s.health_getHealth)
	g.GET("/database", s.health_databaseHealth)

}

func (s *Server) health_getHealth(c echo.Context) error {
	return c.String(200, "OK")
}

func (s *Server) health_databaseHealth(c echo.Context) error {
	return s.db.Health()
}
