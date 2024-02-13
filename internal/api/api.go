package api

import (
	routitle "go-readlist/internal/api/routes/title"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type server struct {
	*chi.Mux
}

func NewAPI() *server {
	return &server{chi.NewRouter()}
}

func (a *server) LoadRoutes() *server {
	a.Mount("/titles", routitle.TitleRoutes())

	return a
}

func (a *server) Run() {
	http.ListenAndServe(":8080", a)
}
