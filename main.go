package main

import (
	"log"
	"net/http"
	"taskmanager/infrastructure/datastore"
	"taskmanager/infrastructure/router"
	"taskmanager/registry"

	"github.com/gorilla/mux"
)

func main() {
	database, err := datastore.NewDB()
	if err != nil {
		log.Fatal("Failed to init DB")
	}

	registry := registry.NewRegistry(database)

	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter = router.NewRouter(muxRouter, registry.NewAppController())

	log.Fatal(http.ListenAndServe("localhost:8080", muxRouter))
}
