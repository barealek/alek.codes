package server

import (
	"fmt"
	"net/http"
	"time"

	"alek.codes/internal/controllers"
	"alek.codes/internal/database"
)

type Server struct {
	port       int
	db         database.Database
	controller controllers.Controller
}

func New() *http.Server {
	// port, _ := strconv.Atoi(os.Getenv("PORT"))
	port := 3000

	db, err := database.New()
	if err != nil {
		panic(fmt.Sprintf("failed to create database: %v\n", err))
	}

	controller := controllers.New(db)

	server := &Server{
		port:       port,
		db:         db,
		controller: controller,
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      server.controller.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
