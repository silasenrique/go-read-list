package routread

import (
	"encoding/json"
	"go-readlist/internal/database"
	"go-readlist/internal/entities"
	"go-readlist/internal/service"
	"net/http"
)

func CreateRead(db database.ICollection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		read := &entities.Read{}

		err := json.NewDecoder(r.Body).Decode(&read)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		service.NewReadCreateService(db).Create(read)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(read)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
