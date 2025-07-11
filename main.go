package main

import (
	"fmt"

	api "github.com/yourusername/savageapp/api"
	database "github.com/yourusername/savageapp/db"
)

func main() {
	database.Initialize()
	db := database.GetDB()
	defer db.Close()
	fmt.Println("Database created or opened successfully!")
	api.StartServer()

}
