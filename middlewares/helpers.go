package middlewares

import (
	"net/http"
)

func GetUserEmailFromRequest(r *http.Request) string {
	if val := r.Context().Value(userEmailKey); val != nil {
		if email, ok := val.(string); ok {
			return email
		}
	}
	return ""
}
