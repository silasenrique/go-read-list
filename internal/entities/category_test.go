package entities

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestCategoryUnmarshalJson_ValidCategory(t *testing.T) {
	scenarios := []struct {
		marshal string
		result  Category
	}{
		{marshal: `{"category": "Manga"}`, result: Category(Manga)},
		{marshal: `{"category": "Book"}`, result: Category(Book)},
		{marshal: `{"category": "WebNovel"}`, result: Category(WebNovel)},
	}

	for _, scenario := range scenarios {
		category := struct {
			Category Category `json:"category"`
		}{}

		err := json.Unmarshal([]byte(scenario.marshal), &category)
		if err != nil {
			t.Errorf("deu erro ao tentar formatar o json, %s", err)
		}

		if category.Category != scenario.result {
			t.Error("deu ruim, ta errado")
		}

	}
}

func TestCategoryUnmarshalJson_InvalidCategory(t *testing.T) {
	t.Run("invalid category", func(t *testing.T) {
		category := `{"category": "Livro"}`
		tempCat := struct {
			Category Category `json:"category"`
		}{}

		err := json.Unmarshal([]byte(category), &tempCat)
		if err != nil && !errors.Is(err, ErrCategoryNotFound) {
			t.Errorf("foi retornado um erro não previsto! err: %s", err)
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		category := `{"category": 30}`
		tempCat := struct {
			Category Category `json:"category"`
		}{}

		err := json.Unmarshal([]byte(category), &tempCat)
		if err == nil {
			t.Errorf("foi retornado um erro não previsto! err: %s", err)
		}
	})

}

func TestCategoryName(t *testing.T) {
	scenarios := []struct {
		name     string
		category Category
	}{
		{name: "Manga", category: Category(Manga)},
		{name: "WebNovel", category: Category(WebNovel)},
		{name: "Book", category: Category(Book)},
	}

	for _, i := range scenarios {
		result := i.category.Name()

		if result != i.name {
			t.Errorf(
				"categoria retornou valor não esperado. Esperado: %s, retornado: %s",
				i.name,
				result,
			)
		}
	}
}
