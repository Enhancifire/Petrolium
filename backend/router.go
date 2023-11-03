package main

import (
	"github.com/Enhancifire/Petrolium/backend/controllers"
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	Home     string
	Petrol   string
	Register string
	Login    string
}

func CreateRouter() *AppRouter {
	router := &AppRouter{
		Home:     "/",
		Petrol:   "/petrol",
		Register: "/register",
		Login:    "/login",
	}
	return router
}

func GetRoutes() *gin.Engine {
	appRoutes := CreateRouter()
	router := gin.Default()

	router.GET(appRoutes.Home, HomeHandler)
	router.GET(appRoutes.Petrol, GetPetrolList)
	router.POST(appRoutes.Register, controllers.Register)
	return router
}
