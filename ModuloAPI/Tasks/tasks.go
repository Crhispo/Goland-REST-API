package Tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type allTasks []task

var tasks allTasks

func init() {

	tasks = allTasks{
		{
			ID:      1,
			Name:    "Task One",
			Content: "Some Content",
		},
	}
}

func GetTasks(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(tasks)
}

func CreateTask(resp http.ResponseWriter, req *http.Request) {
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

func GetTask(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(resp, "Invalid id")
		return
	}
	for _, task := range tasks {
		if task.ID == taskID {
			resp.Header().Set("Content-Type", "application/json")
			json.NewEncoder(resp).Encode(task)
		}
	}
}

func DeleteTask(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(resp, "Invalid id")
		return
	}
	for i, task := range tasks {
		if task.ID == taskID {
			resp.Header().Set("Content-Type", "application/json")
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(resp, "The task with ID %v has been remove succesfully", taskID)
		}
	}
}

func UpdateTask(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	taskID, err := strconv.Atoi(vars["id"])
	var updateTask task
	if err != nil {
		fmt.Fprintf(resp, "Invalid id")
		return
	}

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Fprintf(resp, "Please insert valid data")
	}

	json.Unmarshal(reqBody, &updateTask)

	for i, task := range tasks {
		if task.ID == taskID {
			resp.Header().Set("Content-Type", "application/json")
			tasks = append(tasks[:i], tasks[i+1:]...)
			updateTask.ID = taskID
			tasks = append(tasks, updateTask)
			fmt.Fprintf(resp, "The task with ID %v has been update succesfully", taskID)
		}
	}
}
