package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", ShowAllTasksFunc)
	http.HandleFunc("/add/", AddTaskFunc)
	http.HandleFunc("/add_user", PostAddUser)
	// http.HandleFunc("/admin", HandleAdmin)
	// http.HandleFunc("/change", PostChange)
	// http.HandleFunc("/complete/", CompleteTaskFunc)
	// http.HandleFunc("/completed/", ShowCompleteTasksFunc)
	// http.HandleFunc("/delete/", DeleteTaskFunc)
	// http.HandleFunc("/deleted/", ShowTrashTaskFunc)
	// http.HandleFunc("/edit/", EditTaskFunc)
	// http.HandleFunc("/login", GetLogin)
	// http.HandleFunc("/logout", HandleLogout)
	// http.HandleFunc("/register", PostRegister)
	// http.HandleFunc("/restore/", RestoreTaskFunc)
	// http.HandleFunc("/search/", SearchTaskFunc)
	// http.HandleFunc("/trash/", TrashTaskFunc)
	// http.HandleFunc("/update/", UpdateTaskFunc)

	http.Handle("/static/", http.FileServer(http.Dir("public")))
	log.Print("running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ShowAllTasksFunc show all tasks
func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}

// AddTaskFunc add new task
func AddTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "create task GET doesn't really work"
	} else {
		message = "create task POST"
	}
	w.Write([]byte(message))
}

// PostAddUser add new user
func PostAddUser(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "POST" {
		message = "add a new user POST"
	} else {
		message = "only POST works"
	}
	w.Write([]byte(message))
}
