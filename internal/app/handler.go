package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

type RootHandler func(w http.ResponseWriter, r *http.Request) error

func HandlerWithErrors(fn RootHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			for errToPrint := err; errToPrint != nil; errToPrint = errors.Unwrap(errToPrint) {
				fmt.Printf("%v\n", errToPrint)
			}

			var e APIError
			if errors.As(err, &e) {
				status, message := e.APIError()
				w.WriteHeader(status)
				body, _ := json.Marshal(errorResponse{Message: message})
				w.Write(body)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				body, _ := json.Marshal(errorResponse{Message: "Unexpected error encountered"})
				w.Write(body)
			}
		}
	}
}
