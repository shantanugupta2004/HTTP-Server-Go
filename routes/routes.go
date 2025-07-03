package routes

import (
	"http-server-go/handlers"
	"http-server-go/middlewares"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router{
	r:= mux.NewRouter()
	r.Use(middlewares.LoggingMiddleware)
	
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/hello", handlers.HelloHandler).Methods("GET")
	r.HandleFunc("/greet", handlers.GreetHandler).Methods("GET")
	r.HandleFunc("/greet/{name}", handlers.GreetByNameHandler).Methods("GET")
	r.HandleFunc("/echo", handlers.EchoHandler).Methods("POST")
	r.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUserByIDHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/profile", middlewares.JWTMiddleware(handlers.ProfileHandler)).Methods("GET")
	r.HandleFunc("/upload", middlewares.JWTMiddleware(handlers.UploadFileHandler)).Methods("POST")
	r.HandleFunc("/getFiles", handlers.GetFilesHandler).Methods("GET")
	r.HandleFunc("/getUserFiles", middlewares.JWTMiddleware(handlers.GetFilesbyUserHandler)).Methods("GET")
	r.HandleFunc("/downloadFile", middlewares.JWTMiddleware(handlers.DownloadFileHandler)).Methods("GET")
	r.HandleFunc("/deleteFile", middlewares.JWTMiddleware(handlers.DeleteFileHandler)).Methods("DELETE")
	r.HandleFunc("/generate-share", middlewares.JWTMiddleware(handlers.GenerateShareLinkHandler)).Methods("GET")
	r.HandleFunc("/share/{token}", handlers.ShareDownloadHandler).Methods("GET")
	return r
}