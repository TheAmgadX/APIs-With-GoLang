package api

import (
	"encoding/json"
	"net/http"
)

// "net/http"
// "encoding/json"

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	// Success Code, 200 usally
	Code int
	// Account Balance
	Balance int64
}

// Error Response
type Error struct {
	// Error Code
	Code int
	// Error Message
	Message string
}

func WriteError(w http.ResponseWriter, msg string, statusCode int) {
	response := Error{
		Code:    statusCode,
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	/*
		application/json -> what type of data is in the response
		(JSON in this case)
	*/
	w.WriteHeader(statusCode) // set the status code of the response.

	json.NewEncoder(w).Encode(response) /* convert error struct to json format
	and write it to the ResponseWriter w
	*/
}

// two types of errors may occure.
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		WriteError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		WriteError(w, "An Unexpected Error Occured.", http.StatusInternalServerError)
	}
)
