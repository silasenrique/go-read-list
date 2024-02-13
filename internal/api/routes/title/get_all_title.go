package routitle

import (
	"encoding/json"
	"go-readlist/internal/entities"
	"net/http"
)

func GetAllTitles(w http.ResponseWriter, r *http.Request) {
	titles := []entities.Title{
		{
			Id:       123,
			Name:     "One Piece",
			Category: entities.Manga,
		},
	}

	data, err := json.Marshal(&titles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
