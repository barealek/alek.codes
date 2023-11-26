package controllers

import (
	"net/http"

	"alek.codes/internal/database"

	"github.com/barealek/echowares/echologger"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	RegisterRoutes() http.Handler
}

type service struct {
	db database.Database
}

func New(db database.Database) Controller {
	return &service{
		db: db,
	}
}

func (s *service) RegisterRoutes() http.Handler {
	e := echo.New()

	//Middlewares
	e.Use(echologger.New())

	s.RegisterHealth(e.Group("/health"))

	return e
}
