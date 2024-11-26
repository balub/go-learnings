package main

import (
	// "log"
	"log"
	"net/http"

	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/config"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/database"
	"github.com/balub/go-learnings/simple-gorm-mux-api/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	// initialize environment variables
	config.LoadConfig()

	// initialize database
	DBService := database.New()
	DBService.Migrate()
	defer DBService.Close()

	// initialize mux router
	r := mux.NewRouter()
	s := r.PathPrefix("/v1/api").Subrouter()
	routes.InitializeRoutes(DBService, s)

	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}

}
