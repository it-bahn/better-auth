package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
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

func GenerateJWT(secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
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
		var s = fmt.Sprintf("\t\t%v:%v\n", k, v)
		str += s
	}
	str += end
	return str
}
