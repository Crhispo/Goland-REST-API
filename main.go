package main

import (
	"PackAPI/ModuloAPI/Routes"
	"log"
	"net/http"
)

func main() {

	Routes.RouteIndex()
	Routes.RouteTasks()
	Routes.RouteCreateTasks()
	Routes.RouteTask()
	Routes.RouteUpdateTask()
	Routes.RouteDeleteTask()

	log.Fatal(http.ListenAndServe(":3000", Routes.Router))
}
