package main

import (
	"fmt"

	"github.com/OmidRasouli/amuse-park/database"
	"github.com/OmidRasouli/amuse-park/models"
	"github.com/OmidRasouli/amuse-park/routing"
	"github.com/OmidRasouli/amuse-park/statics"
)

func main() {
	statics.Read()

	database.Initialize(&models.Account{}, &models.Authentication{}, &models.Profile{})

	fmt.Println("Hello there...")

	routing.Initialize().Run(":8080")
}
