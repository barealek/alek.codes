package main

import (
	"alek.codes/internal/server"
)

func main() {
	server := server.New()

	if err := server.ListenAndServe(); err != nil {
		panic("failed to start server")
	}
}
