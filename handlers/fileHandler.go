package handlers

import (
	"encoding/json"
	"fmt"
	"http-server-go/database"
	"http-server-go/middlewares"
	"http-server-go/models"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(10<<20) //max 10MB file
	file, handler, err := r.FormFile("file")
	if err!=nil{
		http.Error(w, "File not given", http.StatusBadRequest)
		return
	}
	defer file.Close()
	os.MkdirAll("uploads", os.ModePerm) //create dir if not exist
	filePath := filepath.Join("uploads", handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Cannot save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	//copying file content
	fileSize, err := io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		return
	}
	userEmail := middlewares.GetUserEmailFromRequest(r)
	var user models.User
	result := database.DB.First(&user, "email = ?", userEmail)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
	newFile := models.File{
		FileName:   handler.Filename,
		FilePath:   filePath,
		FileSize:   fileSize,
		UploadedAt: time.Now(),
		UserID:     user.ID,
	}
	saveErr := database.DB.Create(&newFile).Error
	if saveErr != nil {
		http.Error(w, "Database insert failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File '%s' uploaded successfully!\n", handler.Filename)
}

func GetFilesHandler(w http.ResponseWriter, r *http.Request){
	var files []models.File
	result := database.DB.Find(&files)
	if result.Error!=nil{
		http.Error(w, "Failed to fetch files", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}

func GetFilesbyUserHandler(w http.ResponseWriter, r *http.Request){
	userEmail := middlewares.GetUserEmailFromRequest(r)
	var user models.User
	result:= database.DB.First(&user, "email = ?", userEmail)
	if result.Error != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }
	var files []models.File
	result2 := database.DB.Where("user_id = ?", user.ID).Find(&files)
	if result2.Error!=nil{
		http.Error(w, "Failed to fetch files", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}