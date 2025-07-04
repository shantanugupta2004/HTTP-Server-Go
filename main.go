package main

import (
	"fmt"
	"net/http"
	"http-server-go/routes"
	"http-server-go/database"
	"http-server-go/models"
	"http-server-go/metrics"
)



func main(){
	database.Connect()
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.File{})
	metrics.Init()
	router := routes.SetupRoutes()
	http.Handle("/metrics", metrics.MetricsHandler())
	http.Handle("/", router)
	fmt.Println("Server running on Port 5000")
	http.ListenAndServe(":5000", nil)
}