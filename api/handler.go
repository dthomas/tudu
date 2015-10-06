package api

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/derick/tudu/router"
	"github.com/gorilla/mux"
)

// Handler function for API
func Handler() *mux.Router {
	m := router.API()
	m.Get(router.ShowTodo).Handler(handler(serveTodoShow))
	m.Get(router.IndexTodo).Handler(handler(serveTodoIndex))
	m.Get(router.CreateTodo).Handler(handler(serveTodoCreate))
	m.Get(router.UpdateTodo).Handler(handler(serveTodoUpdate))
	m.Get(router.DeleteTodo).Handler(handler(serveTodoDelete))
	return m
}

type handler func(w http.ResponseWriter, r *http.Request) *AppResponse

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting ", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "application/vnd.api+json")

	if r.Header.Get("Accept") != "application/vnd.api+json" {
		w.WriteHeader(http.StatusNotAcceptable)
		msg := make(map[string][]string)
		msg["header"] = append(msg["Header"], "Accept header must be set to 'application/vnd.api+json'")
		json.NewEncoder(w).Encode(&AppResponse{"Not Acceptable", http.StatusNotAcceptable, nil, msg})
		return
	}

	res := h(w, r)
	w.WriteHeader(res.Code)

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
	}
}
