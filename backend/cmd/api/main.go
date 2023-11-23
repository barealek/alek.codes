package main

import (
	"alek.codes/internal/server"
)

func main() {
	server := server.New()

	err := server.ListenAndServe()

	if err != nil {
		panic("failed to start server")
	}
}
