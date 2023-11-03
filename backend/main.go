package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("Petrolium Backend")

	router := GetRoutes()
	err := router.Run("localhost:8090")
	handleError(err)

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
