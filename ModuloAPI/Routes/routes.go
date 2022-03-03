package Routes

import (
	"PackAPI/ModuloAPI/Tasks"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {

	Router = mux.NewRouter().StrictSlash(false)

}

func RouteIndex() {
	Router.HandleFunc("/", Tasks.IndexRoute).Methods("GET")
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
