package handlers

import (
	"better-auth/internal/models"
	"better-auth/internal/models/users"
	"better-auth/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetErrorFromRequestBody(w http.ResponseWriter, err error) map[string]interface{} {
	var response map[string]interface{}
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

func ValidateReqBody(w http.ResponseWriter, user users.User) map[string]interface{} {
	var res map[string]interface{}

	status := utils.IsValidAddress(user.Address.Street)
	if status == false && user.Address.Street != "" {
		res = models.CreateResponse("failure", "Invalid Address", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return res
	}
	status = utils.IsValidCity(user.Address.City)
	if status == false && user.Address.City != "" {
		res = models.CreateResponse("failure", "Invalid City", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return res
	}
	status = utils.IsValidState(user.Address.State)
	if status == false && user.Address.State != "" {
		res = models.CreateResponse("failure", "Invalid State", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return res
	}
	status = utils.IsValidZip(user.Address.Zip)
	if status == false && user.Address.Zip != "" {
		res = models.CreateResponse("failure", "Invalid Zip", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return res
	}
	status = utils.IsValidCountry(user.Address.Country)
	if status == false && user.Address.Country != "" {
		res = models.CreateResponse("failure", "Invalid Country", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return res
	}
	status = utils.IsValidEmail(user.Email)
	if status == false && user.Email != "" {
		res = models.CreateResponse("failure", "Invalid Email", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return res
	}
	status = utils.IsValidPhone(user.Telephone)
	if status == false && user.Telephone != "" {
		res = models.CreateResponse("failure", "Invalid Phone", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return res
	}

	status = utils.IsValidName(user.LastName)
	if status == false && user.LastName != "" {
		res = models.CreateResponse("failure", "Invalid Last Name", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return res
	}
	res = models.CreateResponse("success", "Valid Request Body", http.StatusOK)
	return res
}

func EnableCors(w *http.ResponseWriter, req *http.Request) {
	if (*req).Header.Get("Origin") != "" {
		(*w).Header().Set("Access-Control-Allow-Origin", (*req).Header.Get("Origin"))
		(*w).Header().Set("Allow-Control-Allow-Origin", (*req).Header.Get("Origin"))
	}
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Allow-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
