package server

import (
	"net/http"

	"github.com/barealek/echowares/echologger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	e.Use(middleware.AddTrailingSlash())
	e.Use(echologger.New())

	s.RegisterHealth(e.Group("/health"))

	return e
}
