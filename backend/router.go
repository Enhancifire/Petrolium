package main

import (
	"github.com/Enhancifire/Petrolium/backend/modules/auth"
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	Home     string
	Petrol   string
	Register string
	Login    string
	PetrolID string
}

func CreateRouter() *AppRouter {
	router := &AppRouter{
		Home:     "/",
		Petrol:   "/petrol",
		Register: "/register",
		Login:    "/login",
		PetrolID: "/petrol/:id",
	}
	return router
}

func GetRoutes() *gin.Engine {
	appRoutes := CreateRouter()
	router := gin.Default()

	router.GET(appRoutes.Home, HomeHandler)
	router.GET(appRoutes.Petrol, GetPetrolList)
	router.GET(appRoutes.PetrolID, GetPetrolInstance)
	router.POST(appRoutes.Petrol, AddPetrolData)
	router.PATCH(appRoutes.PetrolID, UpdatePetrolData)
	router.DELETE(appRoutes.PetrolID, DeletePetrolInstance)
	router.POST(appRoutes.Register, auth.Register)

	return router
}
