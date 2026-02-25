package migration

import (
	"fmt"
	"payment-service/db"
)

func Migration() {
	database := db.Database()

	err := database.AutoMigrate()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
