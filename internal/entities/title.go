package entities

type Title struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
}

func NewTitle(name string, category Category) *Title {
	return &Title{
		Name:     name,
		Category: category,
	}
}
