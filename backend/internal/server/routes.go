package server

import (
	"net/http"

	"github.com/barealek/echowares/echologger"
	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	e.Use(echologger.New())

	// routes.RegisterHealth(e.Group("/health"))
	// TODO: Fix cyclic import error

	return e
}
