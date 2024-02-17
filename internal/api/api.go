package api

import (
	"go-readlist/internal/api/routes/routread"
	"go-readlist/internal/database"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type server struct {
	*chi.Mux
}

func NewAPI() *server {
	return &server{chi.NewRouter()}
}

func (a *server) LoadRoutes(db *database.DB) *server {
	a.Mount("/read", routread.TitleRoutes(db))

	return a
}

func (a *server) Run() {
	http.ListenAndServe(":8080", a)
}
