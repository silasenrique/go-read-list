package main

import (
	"go-readlist/internal/api"
)

func main() {
	api.NewAPI().LoadRoutes().Run()

}
