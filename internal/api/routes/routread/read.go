package routread

import (
	"go-readlist/internal/database"

	"github.com/go-chi/chi/v5"
)

func TitleRoutes(db *database.DB) chi.Router {
	collection := db.GetCollection("readings", "readings")

	r := chi.NewRouter()

	r.Post("/", CreateRead(collection))
	r.Get("/{id}", FindByIdRead(collection))

	return r
}
