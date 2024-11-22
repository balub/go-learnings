package main

import (
	"fmt"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/config"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/database"
	"github.com/gorilla/mux"
)

func main() {
	// initialize environment variables
	config.LoadConfig()

	// initialize database
	DBService := database.New()
	defer DBService.Close()
	fmt.Print(DBService)

	// initialize mux router
	r := mux.NewRouter()

	// initialize router handler function

	http.Handle("/", r)
}
