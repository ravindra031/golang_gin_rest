package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("postgres", "host=<> port=<> user=<> dbname=<> password=<> sslmode=disable")
	//defer database.Close()

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})

	DB = database
}
