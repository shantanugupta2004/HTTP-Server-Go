package handlers

import (
	"encoding/json"
	"net/http"
	"http-server-go/models"
	"http-server-go/database"
	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request){
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err!=nil{
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	result := database.DB.Create(&user)
	if result.Error!=nil{
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request){
	var users []models.User

	result:= database.DB.Find(&users)
	if result.Error!=nil{
		http.Error(w, "Failed to fetch all Users",  http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id := params["id"]
	var user models.User
	result:= database.DB.First(&user, id)
	if result.Error!=nil{
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request){
	params:= mux.Vars(r)
	id:= params["id"]
	var user models.User
	result:= database.DB.First(&user, id)
	if result.Error!=nil{
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	var updatedData models.User
	err:= json.NewDecoder(r.Body).Decode(&updatedData);
	if err!=nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
	}
	user.Name = updatedData.Name
	user.Email = updatedData.Email
	database.DB.Save(&user)
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id:=params["id"]
	result:= database.DB.Delete(&models.User{}, id)
	if result.Error != nil {
        http.Error(w, "Failed to delete user", http.StatusInternalServerError)
        return
    }
	w.WriteHeader(http.StatusNoContent)
}