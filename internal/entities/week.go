package entities

import (
	"encoding/json"
	"errors"
)

var (
	ErrDayNotFound = errors.New("dia invalido")
)

type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w *WeekDay) toWeekDay(day string) error {
	switch day {
	case "Sunday":
		*w = Sunday
	case "Monday":
		*w = Monday
	case "Tuesday":
		*w = Tuesday
	case "Wednesday":
		*w = Wednesday
	case "Thursday":
		*w = Thursday
	case "Friday":
		*w = Friday
	case "Saturday":
		*w = Saturday
	default:
		return ErrDayNotFound
	}

	return nil
}

func (w *WeekDay) Name() string {
	switch *w {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	default:
		return "Saturday"
	}
}

func (w *WeekDay) UnmarshalJSON(data []byte) error {
	var day string

	err := json.Unmarshal(data, &day)
	if err != nil {
		return err
	}

	return w.toWeekDay(day)
}

func (w *WeekDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.Name())
}
