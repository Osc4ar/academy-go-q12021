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
	reader := datastore.NewDB()
	defer datastore.CloseDB()

	registry := registry.NewRegistry(reader)

	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter = router.NewRouter(muxRouter, registry.NewAppController())

	log.Fatal(http.ListenAndServe("localhost:8080", muxRouter))
}
