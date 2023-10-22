package models

import (
	"log"

	"gorm.io/gorm"
)

func InitModels(conn *gorm.DB) {
	err := conn.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

}
