package router

import (
	"fmt"
	"net/http"
	"strconv"
	"taskmanager/interface/controllers"

	"github.com/gorilla/mux"
)

// NewRouter defines the endpoints of the Application
func NewRouter(muxRouter *mux.Router, c controllers.AppController) *mux.Router {
	muxRouter.HandleFunc("/tasks", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("content-type", "application/json")

		if err := c.GetTasks(rw); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, "Could not retrieve tasks: %v", err)
		}
	})

	muxRouter.HandleFunc("/tasks/{id}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		rw.Header().Add("content-type", "application/json")

		if err := c.GetTask(uint(id), rw); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(rw, "Could not retrieve task: %v", err)
		}
	})

	muxRouter.HandleFunc("/cocurrent-task", func(rw http.ResponseWriter, r *http.Request) {
		objectsType := r.URL.Query().Get("type")
		items := r.URL.Query().Get("items")
		itemsPerWorker := r.URL.Query().Get("items_per_workers")

		itemsInt, err := strconv.Atoi(items)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		itemsPerWorkerInt, err := strconv.Atoi(itemsPerWorker)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		if objectsType != "odd" && objectsType != "even" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		rw.Header().Add("content-type", "application/json")

		c.GetTaskConcurrently(objectsType == "even", itemsInt, itemsPerWorkerInt, rw)
	})

	return muxRouter
}
