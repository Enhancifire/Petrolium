package main

import (
	"fmt"

	"github.com/Enhancifire/Petrolium/backend/modules/auth"
	"github.com/Enhancifire/Petrolium/backend/modules/petrol"
	"github.com/Enhancifire/Petrolium/backend/modules/storage"
)

func SetupDatabase() {
	fmt.Println("Setting up database")
	storage.ConnectToDatabase()
	petrolModels := petrol.ReturnModels()
	userModels := auth.ReturnModels()
	models := append(petrolModels, userModels...)

	for _, item := range models {
		fmt.Printf("Migrating %T\n", item)
		storage.Migrate(item)
	}

}
