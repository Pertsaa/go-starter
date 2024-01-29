package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type server struct {
	router *chi.Mux
}

func newServer(r *chi.Mux) *server {
	s := &server{
		router: r,
	}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func Run() error {
	fmt.Println("Running server app...")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// r.Use(middleware.Timeout(60 * time.Second))

	s := newServer(r)

	http.ListenAndServe(":8080", s)

	return nil
}
