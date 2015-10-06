package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJSON(w http.ResponseWriter, v interface{}) *AppResponse {
	w.Header().Set("content-type", "application/json; charset=utf-8")

	data, err := json.Marshal(v)
	if err != nil {
		errString := fmt.Sprintf("%s", err)
		return &AppResponse{errString, 500, nil, nil}
	}

	_, err = w.Write(data)
	if err != nil {
		errString := fmt.Sprintf("%s", err)
		return &AppResponse{errString, 500, nil, nil}
	}

	return nil
}
