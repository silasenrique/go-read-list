package entities

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestWeekDayUnmarshalJson_ValidDays(t *testing.T) {
	scenarios := []struct {
		marshal string
		day     WeekDay
	}{
		{marshal: `{"day": "Sunday"}`, day: WeekDay(Sunday)},
		{marshal: `{"day": "Saturday"}`, day: WeekDay(Saturday)},
		{marshal: `{"day": "Friday"}`, day: WeekDay(Friday)},
		{marshal: `{"day": "Monday"}`, day: WeekDay(Monday)},
		{marshal: `{"day": "Tuesday"}`, day: WeekDay(Tuesday)},
		{marshal: `{"day": "Wednesday"}`, day: WeekDay(Wednesday)},
		{marshal: `{"day": "Thursday"}`, day: WeekDay(Thursday)},
	}

	for _, scenario := range scenarios {
		week := struct {
			Day WeekDay `json:"day"`
		}{}

		err := json.Unmarshal([]byte(scenario.marshal), &week)
		if err != nil {
			t.Errorf("Houve um erro não previsto! Err: %s", err)
		}

		if week.Day != scenario.day {
			t.Errorf(
				"era esperado %s, foi retornado %s",
				scenario.day.Name(),
				week.Day.Name(),
			)
		}
	}
}

func TestWeekDayUnmarshalJson_InvalidDay(t *testing.T) {
	t.Run("invalid day", func(t *testing.T) {
		weekDay := `{"day": "Sabado"}`
		week := struct {
			Day WeekDay `json:"day"`
		}{}

		err := json.Unmarshal([]byte(weekDay), &week)
		if err != nil && !errors.Is(err, ErrDayNotFound) {
			t.Errorf("foi retornado um erro não previsto! err: %s", err)
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		weekDay := `{"day": 20}`
		week := struct {
			Day WeekDay `json:"day"`
		}{}

		err := json.Unmarshal([]byte(weekDay), &week)
		if err == nil {
			t.Errorf("foi retornado um erro não previsto! err: %s", err)
		}
	})

}

func TestWeekDayName(t *testing.T) {

	scenarios := []struct {
		name string
		day  WeekDay
	}{
		{name: "Sunday", day: WeekDay(Sunday)},
		{name: "Saturday", day: WeekDay(Saturday)},
		{name: "Friday", day: WeekDay(Friday)},
		{name: "Monday", day: WeekDay(Monday)},
		{name: "Tuesday", day: WeekDay(Tuesday)},
		{name: "Wednesday", day: WeekDay(Wednesday)},
		{name: "Thursday", day: WeekDay(Thursday)},
	}

	for _, scenario := range scenarios {
		result := scenario.day.Name()
		if result != scenario.name {
			t.Errorf("nome invalido. era esperado %s, retornado %s", scenario.name, result)
		}
	}
}
