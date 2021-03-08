package router

import (
	"net/http"
	"strconv"
	"taskmanager/interface/controllers"

	"github.com/gorilla/mux"
)

// NewRouter defines the endpoints of the Application
func NewRouter(muxRouter *mux.Router, c controllers.AppController) *mux.Router {
	muxRouter.HandleFunc("/tasks", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("content-type", "application/json")
		c.GetTasks(rw)
	})

	muxRouter.HandleFunc("/tasks/{id}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		rw.Header().Add("content-type", "application/json")
		c.GetTask(uint(id), rw)
	})

	return muxRouter
}
