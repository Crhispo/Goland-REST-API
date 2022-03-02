package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}

func getTasks(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(tasks)
}

func createTask(resp http.ResponseWriter, req *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Fprintf(resp, "Insert a valid Task")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1

	tasks = append(tasks, newTask)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusCreated)

	json.NewEncoder(resp).Encode(newTask)

}

func indexRoute(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Welcome to my API")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks)
	log.Fatal(http.ListenAndServe(":3000", router))
}
