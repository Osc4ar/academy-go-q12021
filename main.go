package main

import (
	"log"
	"net/http"
	"taskmanager/infrastructure/datastore"
	"taskmanager/infrastructure/router"
	"taskmanager/registry"
)

func main() {
	reader := datastore.NewDB()

	registry := registry.NewRegistry(reader)

	router.NewRouter(registry.NewAppController())

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
