package Routes

import (
	"PackAPI/ModuloAPI/Tasks"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexRoute(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Welcome to my API")
}

var Router *mux.Router

func init() {

	Router = mux.NewRouter().StrictSlash(true)

}

func RouteIndex() {
	Router.HandleFunc("/", indexRoute)
}

func RouteTasks() {
	Router.HandleFunc("/tasks", Tasks.GetTasks).Methods("GET")
}

func RouteCreateTasks() {
	Router.HandleFunc("/tasks", Tasks.CreateTask).Methods("POST")
}

func RouteTask() {
	Router.HandleFunc("/tasks/{id}", Tasks.GetTask).Methods("GET")
}

func RouteDeleteTask() {
	Router.HandleFunc("/tasks/{id}", Tasks.DeleteTask).Methods("DELETE")
}

func RouteUpdateTask() {
	Router.HandleFunc("/tasks/{id}", Tasks.UpdateTask).Methods("PUT")
}
