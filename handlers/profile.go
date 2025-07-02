package handlers

import(
	"net/http"
	"encoding/json"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Access granted to protected route",
	})
}