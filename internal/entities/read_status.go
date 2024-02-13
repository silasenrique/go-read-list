package entities

import (
	"encoding/json"
	"errors"
)

type ReadingState int

const (
	NotStarted ReadingState = iota
	InProgress
	Completed
)

func (r *ReadingState) toReadingState(state string) error {
	switch state {
	case "NotStarted":
		*r = NotStarted
	case "InProgress":
		*r = InProgress
	case "Completed":
		*r = Completed
	default:
		return errors.New("estado de leitura não é valido")
	}

	return nil
}

func (r *ReadingState) Name() string {
	switch *r {
	case NotStarted:
		return "NotStarted"
	case InProgress:
		return "InProgress"
	default:
		return "Completed"
	}
}

func (r *ReadingState) UnmarshalJSON(data []byte) error {
	var state string

	err := json.Unmarshal(data, &state)
	if err != nil {
		return nil
	}

	return r.toReadingState(state)
}

func (r *ReadingState) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Name())
}
