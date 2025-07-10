package main

import (
	"fmt"
	"net/http"
	"http-server-go/routes"
	"http-server-go/database"
	"http-server-go/models"
	"http-server-go/metrics"
	"http-server-go/middlewares"
)



func main(){
	database.Connect()
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.File{})
	metrics.Init()
	router := routes.SetupRoutes()
	corsWrappedRouter := middlewares.CORSMiddleware(router)
	http.Handle("/metrics", metrics.MetricsHandler())
	http.Handle("/", corsWrappedRouter)
	fmt.Println("Server running on Port 5000")
	http.ListenAndServe(":5000", nil)
}