package main

import (
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	Home   string
	Petrol string
}

func CreateRouter() *AppRouter {
	router := &AppRouter{
		Home:   "/",
		Petrol: "/petrol",
	}
	return router
}

func GetRoutes() *gin.Engine {
	appRoutes := CreateRouter()
	router := gin.Default()

	router.GET(appRoutes.Home, HomeHandler)
	router.GET(appRoutes.Petrol, GetPetrolList)
	return router
}
