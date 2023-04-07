package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID          string `json:"id"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Description string `json:"product_detail"`
	Price       string `json:"price"`
	Category    string `json:"category"`
}

var tasks []Tasks

func allTasks() {
	task := Tasks{
		ID:          "1",
		Image:       "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
		Title:       "New Project",
		Description: "You must lead the project",
		Category:    "men's clothing",
		Price:       "100"}

	tasks = append(tasks, task)

	task1 := Tasks{
		ID:          "1",
		Image:       "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
		Title:       "New Project",
		Description: "You must lead the project",
		Category:    "men's clothing",
		Price:       "100"}

	tasks = append(tasks, task1)

	task2 := Tasks{
		ID:          "1",
		Image:       "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
		Title:       "New Project",
		Description: "You must lead the project",
		Category:    "men's clothing",
		Price:       "100"}

	tasks = append(tasks, task2)

	task3 := Tasks{
		ID:          "1",
		Image:       "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
		Title:       "New Project",
		Description: "You must lead the project",
		Category:    "men's clothing",
		Price:       "100"}

	tasks = append(tasks, task3)

	fmt.Println(tasks)

}

func gettasks(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "gettasks is hit")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func handleRoute() {
	router := mux.NewRouter()
	//router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", gettasks).Methods("GET")
	// router.HandleFunc("/gettask/{id}", gettask).Methods("GET")
	// router.HandleFunc("/create", createTask).Methods("POST")
	// router.HandleFunc("/delete{id}", deleteTask).Methods("DELETE")
	// router.HandleFunc("/update", updateTask).Queries("id", "{id}").Methods("PUT")
	http.ListenAndServe(":8080", router)
}

func main() {

	allTasks()
	fmt.Println("Hello World")
	handleRoute()

}
