package main

import (
	"log"
	"net/http"
	"os"
	"taskmanager/infrastructure/client"
	"taskmanager/infrastructure/datastore"
	"taskmanager/infrastructure/router"
	"taskmanager/registry"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load env file")
	}

	token := os.Getenv("API_TOKEN")
	projectId := os.Getenv("PROJECT_ID")
	apiClient := client.NewApiClient(token, projectId)

	database, err := datastore.NewDB()
	if err != nil {
		log.Fatal("Failed to init DB")
	}

	registry := registry.NewRegistry(database, apiClient)

	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter = router.NewRouter(muxRouter, registry.NewAppController())

	log.Fatal(http.ListenAndServe("localhost:8080", muxRouter))
}
