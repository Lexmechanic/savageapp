package api

import (
	"fmt"
	"log"
	"net/http"

	encounter "github.com/yourusername/savageapp/encounter"
)

func StartServer() {
	// Set up the HTTP server and routes
	BestiaryAPIRoutes()
	SkillAPIRoutes()
	http.HandleFunc("/ws", encounter.HandleNewConnection)

	fmt.Println("API server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
