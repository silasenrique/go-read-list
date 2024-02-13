package entities

import "encoding/json"

type Read struct {
	Title
	ReadInfo
}

func (r *Read) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Id       int      `json:"id"`
			Name     string   `json:"name"`
			Category Category `json:"category"`
			ReadInfo ReadInfo `json:"readInfo"`
		}{
			Id:       r.Id,
			Name:     r.Name,
			Category: r.Category,
			ReadInfo: r.ReadInfo,
		},
	)
}
