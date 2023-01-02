package main

import (
	"github.com/ridhompra/models"
	"github.com/ridhompra/routers"
)

func main() {
	models.DbInit()
	routers.Routers()
}
