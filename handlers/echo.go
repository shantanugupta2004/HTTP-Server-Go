package handlers

import(
	"fmt"
	"net/http"
	"http-server-go/models"
	"encoding/json"
)

func EchoHandler(w http.ResponseWriter, r *http.Request){
	if r.Method!="POST"{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var reqbody models.EchoRequest;
	err := json.NewDecoder(r.Body).Decode(&reqbody)
	if err!=nil{
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	response:= map[string]string{
		"message": fmt.Sprintf("Hello %s!", reqbody.Name),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}