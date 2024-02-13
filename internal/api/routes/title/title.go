package routitle

import "github.com/go-chi/chi/v5"

func TitleRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", GetAllTitles)

	return r
}
