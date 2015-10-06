package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bitbucket.org/derick/tudu/datastore"

	"github.com/gorilla/mux"
)

func serveTodoShow(w http.ResponseWriter, r *http.Request) *AppResponse {
	id, err := strconv.Atoi(mux.Vars(r)["ID"])
	if err != nil {
		return &AppResponse{"Invalid ID", http.StatusBadRequest, nil, nil} // Send 400 Bad Request
	}

	todo := &datastore.Todo{}
	err = todo.Get(id)
	if err != nil {
		return &AppResponse{"Resource Not found", http.StatusNotFound, nil, nil} // Send 404 Not Found
	}
	resp := &AppResponse{
		Message: "Success",
		Code:    http.StatusOK,
		Data:    todo,
	}
	return resp // Send 200 OK
}

func serveTodoIndex(w http.ResponseWriter, r *http.Request) *AppResponse {
	todos, err := datastore.GetAllTodo()
	if err != nil {
		return &AppResponse{"Internal Server Error", http.StatusInternalServerError, nil, nil}
	}
	resp := &AppResponse{
		Message: "Success",
		Code:    http.StatusOK,
		Data:    todos,
	}
	return resp
}

func serveTodoCreate(w http.ResponseWriter, r *http.Request) *AppResponse {
	todo := new(datastore.Todo)
	err := json.NewDecoder(r.Body).Decode(&AppRequest{todo})
	if err != nil {
		return &AppResponse{"Bad Request", http.StatusBadRequest, nil, nil}
	}

	errors, success := todo.Validate()

	if success != true {
		return &AppResponse{"Unprocessable Entity", 422, todo, errors}
	}

	err = todo.Save()

	if err != nil {
		return &AppResponse{err.Error(), 422, nil, nil}
	}

	resp := &AppResponse{"Success", http.StatusCreated, todo, nil}
	return resp
}

func serveTodoUpdate(w http.ResponseWriter, r *http.Request) *AppResponse {
	todo := new(datastore.Todo)
	err := json.NewDecoder(r.Body).Decode(&AppRequest{todo})
	if err != nil {
		return &AppResponse{"Bad Request", http.StatusBadRequest, todo, nil}
	}

	id, err := strconv.Atoi(mux.Vars(r)["ID"])
	if err != nil {
		return &AppResponse{"Invalid ID", http.StatusBadRequest, nil, nil} // Send 400 Bad Request
	}

	todo.ID = id

	errors, success := todo.Validate()

	if success != true {
		return &AppResponse{"Unprocessable Entity", 422, todo, errors}
	}

	err = todo.Update()
	if err != nil {
		return &AppResponse{err.Error(), 422, todo, nil}
	}

	resp := &AppResponse{"Success", http.StatusOK, todo, nil}
	return resp
}

func serveTodoDelete(w http.ResponseWriter, r *http.Request) *AppResponse {
	id, err := strconv.Atoi(mux.Vars(r)["ID"])
	if err != nil {
		return &AppResponse{"Invalid ID", http.StatusBadRequest, nil, nil} // Send 400 Bad Request
	}

	todo := &datastore.Todo{}
	err = todo.Delete(id)

	if err != nil {
		return &AppResponse{"Unprocessable Entity", 422, nil, nil} // Send 422 Not Found
	}
	resp := &AppResponse{
		Message: "Success",
		Code:    http.StatusOK,
	}
	return resp // Send 200 OK
}
