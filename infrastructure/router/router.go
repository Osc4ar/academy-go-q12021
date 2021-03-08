package router

import (
	"net/http"
	"taskmanager/interface/controllers"
)

// NewRouter defines the endpoints of the Application
func NewRouter(c controllers.AppController) {
	http.HandleFunc("/alltasks", func(rw http.ResponseWriter, r *http.Request) {
		c.GetTasks(rw)
	})
}
