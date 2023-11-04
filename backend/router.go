package main

import (
	"github.com/Enhancifire/Petrolium/backend/modules/auth"
	"github.com/Enhancifire/Petrolium/backend/modules/middleware"
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	Home     string
	Petrol   string
	Register string
	Login    string
	Users    string
}

func CreateRouter() *AppRouter {
	router := &AppRouter{
		Home:     "/",
		Petrol:   "/petrol",
		Register: "/register",
		Login:    "/login",
		Users:    "/users",
	}
	return router
}

func GetRoutes() *gin.Engine {
	appRoutes := CreateRouter()
	router := gin.Default()

	router.GET(appRoutes.Home, HomeHandler)
	router.POST(appRoutes.Register, auth.Register)
	router.POST(appRoutes.Login, auth.Login)

	users := router.Group(appRoutes.Users)
	users.Use(middleware.JWTMiddleware())
	users.GET("", auth.CurrentUser)

	petrol := router.Group(appRoutes.Petrol)
	petrol.Use(middleware.JWTMiddleware())
	petrol.GET("", GetPetrolList)
	router.GET("/:id", GetPetrolInstance)
	router.POST("", AddPetrolData)
	router.PATCH("/:id", UpdatePetrolData)
	router.DELETE("/:id", DeletePetrolInstance)

	return router
}
