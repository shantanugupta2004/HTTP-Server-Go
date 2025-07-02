package database

import(
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	dsn:= "host=localhost user=postgres password=postgres dbname=go_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic("Failed to connect to DB: "+ err.Error())
	}
	fmt.Println("Connected to Postgres DB")
	DB = db
}