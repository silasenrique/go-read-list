package routread

import (
	"encoding/json"
	"go-readlist/internal/database"
	"go-readlist/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func FindByIdRead(db database.ICollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		result, err := service.NewReadCreateService(db).FindById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
