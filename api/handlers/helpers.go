package handlers

import (
	"better-auth/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func ParseErrorFromRequest(w http.ResponseWriter, response map[string]interface{}, err error) map[string]interface{} {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxError):
		log.Printf("Error: %s", err.Error())
		msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		response = models.CreateResponse("failure", msg, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, io.ErrUnexpectedEOF):
		log.Printf("Error: %s", err.Error())
		msg := fmt.Sprintf("Request body contains badly-formed JSON")
		response = models.CreateResponse("failure", msg, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	case errors.As(err, &unmarshalTypeError):
		log.Printf("Error: %s", err.Error())
		msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		response = models.CreateResponse("failure", msg, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		log.Printf("Error: %s", err.Error())
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
		response = models.CreateResponse("failure", msg, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, io.EOF):
		log.Printf("Error: %s", err.Error())
		msg := "Request body must not be empty"
		response = models.CreateResponse("failure", msg, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	case err.Error() == "http: request body too large":
		log.Printf("Error: %s", err.Error())
		msg := "Request body must not be larger than 1MB"
		response = models.CreateResponse("failure", msg, http.StatusRequestEntityTooLarge)
		w.WriteHeader(http.StatusRequestEntityTooLarge)
	default:
		log.Printf("Error: %s", err.Error())
		msg := "Oops! Something went wrong"
		response = models.CreateResponse("failure", msg, http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
	}

	return response
}
