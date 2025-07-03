package database

import(
	"fmt"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect(){
	_ = godotenv.Load()
	dsn:= fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
					 os.Getenv("DB_HOST"),
					 os.Getenv("DB_USER"),
					 os.Getenv("DB_PASSWORD"),
					 os.Getenv("DB_NAME"),
					 os.Getenv("DB_PORT"),					 
)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic("Failed to connect to DB: "+ err.Error())
	}
	fmt.Println("Connected to Postgres DB")
	DB = db
}