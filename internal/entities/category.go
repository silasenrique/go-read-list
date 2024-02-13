package entities

import (
	"encoding/json"
	"errors"
)

var (
	ErrCategoryNotFound = errors.New("categoria inválida")
)

const (
	Manga Category = iota
	Book
	WebNovel
)

type Category int

func (c *Category) toCategory(category string) error {
	switch category {
	case "Manga":
		*c = Manga
	case "Book":
		*c = Book
	case "WebNovel":
		*c = WebNovel
	default:
		return ErrCategoryNotFound
	}

	return nil
}

func (c *Category) Name() string {
	switch *c {
	case Manga:
		return "Manga"
	case Book:
		return "Book"
	default:
		return "WebNovel"
	}
}

func (c *Category) UnmarshalJSON(data []byte) error {
	var category string

	err := json.Unmarshal(data, &category)
	if err != nil {
		return err
	}

	return c.toCategory(category)
}

func (c *Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name())
}
