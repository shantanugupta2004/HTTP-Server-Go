package handlers

import (
	"encoding/json"
	"fmt"
	"http-server-go/database"
	"http-server-go/middlewares"
	"http-server-go/models"
	"github.com/google/uuid"
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
		IsShared:   false,
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

func DownloadFileHandler(w http.ResponseWriter, r *http.Request){
	fileName := r.URL.Query().Get("name")
	if fileName==""{
		http.Error(w, "Missing file name", http.StatusBadRequest)
		return
	}
	filePath := filepath.Join("uploads", fileName)
	http.ServeFile(w, r, filePath)
}

func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("name")
	if fileName == "" {
		http.Error(w, "Missing file name", http.StatusBadRequest)
		return
	}

	userEmail := middlewares.GetUserEmailFromRequest(r)
	var user models.User
	database.DB.First(&user, "email = ?", userEmail)

	var file models.File
	result := database.DB.First(&file, "file_name = ? AND user_id = ?", fileName, user.ID)
	if result.Error != nil {
		http.Error(w, "File not found or unauthorized", http.StatusNotFound)
		return
	}

	err := os.Remove(file.FilePath)
	if err != nil {
		http.Error(w, "Failed to delete file", http.StatusInternalServerError)
		return
	}
	database.DB.Delete(&file)
	fmt.Fprintf(w, "File '%s' deleted successfully", fileName)
}

func GenerateShareLinkHandler(w http.ResponseWriter, r *http.Request){
	fileName := r.URL.Query().Get("name")
	if fileName==""{
		http.Error(w, "File name missing", http.StatusBadRequest)
		return
	}
	userEmail := middlewares.GetUserEmailFromRequest(r)
	var user models.User
	err:= database.DB.First(&user, "email = ?", userEmail).Error
	if err!=nil{
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
	var file models.File
	err2 := database.DB.First(&file, "file_name = ? AND user_id = ?", fileName, user.ID).Error
	if err2!=nil{
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	file.ShareToken = uuid.New().String()
	file.IsShared = true
	database.DB.Save(&file)
	shareURL := fmt.Sprintf("http://localhost:5000/share/%s", file.ShareToken)
	json.NewEncoder(w).Encode(map[string]string{"url": shareURL})
}

func ShareDownloadHandler(w http.ResponseWriter, r *http.Request){
	token := r.URL.Path[len("/share/"):]
	if token==""{
		http.Error(w, "Missing share token", http.StatusBadRequest)
		return
	}
	var file models.File
	err:= database.DB.First(&file, "share_token = ? AND is_shared = ?", token, true).Error
	if err!=nil{
		http.Error(w, "Invalid or expired share link", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, file.FilePath)
}
