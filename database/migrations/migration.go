package migrations

import (
	"crud-fiber/database"
	"crud-fiber/models/entity"
	"fmt"
	"log"
)

func Migrate() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err, err.Error())
	}
	fmt.Println("🚀 Successfully run migrations")
}