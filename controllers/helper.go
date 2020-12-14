package controllers

import (
	"encoding/json"
	"net/http"
)

//Response common reponse
type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

//ProcessResponse process http response
func ProcessResponse(w http.ResponseWriter, message string, statusCode int, data interface{}) {
	response := Response{}
	response.Status = "SUCCESS"
	response.Message = message
	response.StatusCode = statusCode
	response.Data = data
	if statusCode == 500 {
		response.Message = "An unexpected error has occured"
	}
	if statusCode != 200 {
		response.Status = "FAILED"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}
