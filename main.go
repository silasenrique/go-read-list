package main

import (
	"go-readlist/internal/api"
	"go-readlist/internal/database"
)

func main() {
	db := database.NewConnection("mongodb://localhost:27017")

	api.NewAPI().LoadRoutes(db).Run()
}
