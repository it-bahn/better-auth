package utils

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"regexp"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func StringToBase64(any string) string {
	return base64.StdEncoding.EncodeToString([]byte(any))
}

func GenerateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	log.Printf("Random String: %v", string(b))
	return string(b)
}

func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func MapToString(v map[string]interface{}) string {
	var str string
	var start = "{\n"
	var end = "\n}"
	str += start
	for k, v := range v {
		res, ok := v.(map[string]interface{})
		if ok {
			str += MapToString(res)
		}
		var s = fmt.Sprintf("\t\t%v:%v\n", k, v)
		str += s
	}
	str += end
	return str
}
func IsValidEmail(email string) bool {
	regexEmail := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	// Verify Input email with regexEmail
	matched, err := regexp.MatchString(regexEmail, email)
	if err != nil {
		log.Println("Email is not valid", err)
	}
	return matched
}

func IsValidName(name string) bool {
	regexName := `^[a-zA-Z]+$`
	// Verify Input email with regexEmail
	matched, err := regexp.MatchString(regexName, name)
	if err != nil {
		log.Println("Name is not valid", err)
	}
	return matched
}
func IsValidPhone(phone string) bool {
	regexPhone := `^[0-9]{10}$`
	// Verify Input email with regexEmail
	matched, err := regexp.MatchString(regexPhone, phone)
	if err != nil {
		log.Println("Phone is not valid", err)
	}
	return matched
}
func IsValidAddress(address string) bool {
	regexAddress := `^[a-zA-Z0-9\s,'-]{1,}$`
	// Verify Input email with regexEmail
	matched, err := regexp.MatchString(regexAddress, address)
	if err != nil {
		log.Println("Address is not valid", err)
	}
	return matched
}
func IsValidCity(city string) bool {
	regexCity := `^[a-zA-Z]+$`
	// Verify Input email with regexEmail
	matched, err := regexp.MatchString(regexCity, city)
	if err != nil {
		log.Println("City is not valid", err)
	}
	return matched
}
func IsValidState(state string) bool {
	regexState := `^[a-zA-Z]+$`
	// Verify Input email with regexEmail
	matched, err := regexp.MatchString(regexState, state)
	if err != nil {
		log.Println("State is not valid", err)
	}
	return matched
}
func IsValidZip(zip string) bool {
	regexZip := `^[0-9]{5}$`
	// Verify Input email with regexEmail
	matched, err := regexp.MatchString(regexZip, zip)
	if err != nil {
		log.Println("Zip is not valid", err)
	}
	return matched
}
func IsValidCountry(country string) bool {
	regexCountry := `^[a-zA-Z]+$`
	// Verify Input email with regexEmail
	matched, err := regexp.MatchString(regexCountry, country)
	if err != nil {
		log.Println("Country is not valid", err)
	}
	return matched
}
