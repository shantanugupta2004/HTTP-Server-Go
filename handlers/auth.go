package handlers

import(
	"encoding/json"
	"net/http"
	"time"
	"http-server-go/models"
	"http-server-go/database"
	"http-server-go/authutils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("shantanu@2004")

// type Claims struct{
// 	Email string `json:"email"`
// 	jwt.RegisteredClaims
// }

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	hashedPassword, _:= bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	result:= database.DB.Create(&user)
	if result.Error!=nil{
		http.Error(w, "User already exists or invalid", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	var input models.User
	json.NewDecoder(r.Body).Decode(&input)
	var user models.User
	result:= database.DB.Where("email = ?", input.Email).First(&user)
	if result.Error!=nil{
		http.Error(w, "Invaild email or password", http.StatusUnauthorized)
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err!=nil{
		http.Error(w, "Invaild email or password", http.StatusUnauthorized)
		return
	}
	expTime := time.Now().Add(1*time.Hour)
	claims := &authutils.Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	 if err != nil {
        http.Error(w, "Could not login", http.StatusInternalServerError)
        return
    }
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

