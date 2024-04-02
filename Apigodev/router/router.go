package router

import (
	"net/http"

	"Apigodev/tasks"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", tasks.GetTasks).Methods("GET")
	router.HandleFunc("/task", tasks.AddTask).Methods("POST")
	router.HandleFunc("/task/{id}", tasks.RemoveTask).Methods("DELETE")
	router.HandleFunc("/task/{id}", tasks.UpdateTask).Methods("PUT")
	return router
}

func StartServer() {
	http.ListenAndServe(":8000", SetupRoutes())
}
