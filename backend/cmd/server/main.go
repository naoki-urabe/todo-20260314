package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	app := NewApp()
	if err := app.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
