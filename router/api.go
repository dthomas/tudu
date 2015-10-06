package router

import "github.com/gorilla/mux"

// API Router
func API() *mux.Router {
	m := mux.NewRouter()
	m.Path("/todos").Methods("GET").Name(IndexTodo)
	m.Path("/todos").Methods("POST").Name(CreateTodo)
	m.Path("/todos/{ID:[0-9]+}").Methods("GET").Name(ShowTodo)
	m.Path("/todos/{ID:[0-9]+}").Methods("PATCH").Name(UpdateTodo)
	m.Path("/todos/{ID:[0-9]+}").Methods("DELETE").Name(DeleteTodo)
	return m
}
