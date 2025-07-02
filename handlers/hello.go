package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r * http.Request){
	fmt.Fprintf(w, "Welcome to Go Server!")
}

func HelloHandler(w http.ResponseWriter, r *http.Request){
	response := map[string]string{"message": "Hello World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GreetHandler(w http.ResponseWriter, r *http.Request){
	name:= r.URL.Query().Get("name")
	if name == ""{
		name = "Guest"
	}

	response:= map[string]string{
		"message": fmt.Sprintf("Hello %s!", name),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GreetByNameHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	name := vars["name"]
	if name == ""{
		name = "Guest"
	}
	response := map[string]string{
		"message": fmt.Sprintf("Hello %s!", name),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

