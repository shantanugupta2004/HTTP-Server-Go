package main

import (
	"fmt"
	"net/http"
	"http-server-go/routes"
	"http-server-go/database"
	"http-server-go/models"
)



func main(){
	database.Connect()
	database.DB.AutoMigrate(&models.User{})
	router := routes.SetupRoutes()
	fmt.Println("Server running on Port 5000")
	http.ListenAndServe(":5000", router)
}