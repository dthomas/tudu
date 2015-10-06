package api

// AppResponse Standard JSON Response
type AppResponse struct {
	Message string              `json:"message"`
	Code    int                 `json:"code"`
	Data    interface{}         `json:"data,omitempty"`
	Errors  map[string][]string `json:"errors,omitempty"`
}

// AppRequest Standard JSON Request
type AppRequest struct {
	Data interface{} `json:"data, omitempty"`
}
