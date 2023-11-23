package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"alek.codes/internal/database"
)

type Server struct {
	port int
	db   database.Database
}

func New() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	db, err := database.New()
	if err != nil {
		panic(fmt.Sprintf("failed to create database: %v\n", err))
	}

	server := &Server{
		port: port,
		db:   db,
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      server.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
