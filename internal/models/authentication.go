package models

import (
	"better-auth/internal/db"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthEmail struct {
	Email    string `json:"email,omitempty,required" bson:"email,omitempty,required"`
	Password string `json:"password,omitempty,required" bson:"password,omitempty,required"`
}

type AuthenticationUserName struct {
	UserName string `json:"user_name,omitempty,required" bson:"user_name,omitempty,required"`
	Password string `json:"password,omitempty,required" bson:"password,omitempty,required"`
}

func (a *AuthEmail) IsEmpty() bool {
	return a.Email == "" || a.Password == ""
}

func (a *AuthenticationUserName) IsEmpty() bool {
	return a.UserName == "" || a.Password == ""
}
func (A *AuthEmail) Login(ctx *context.Context) map[string]interface{} {
	filter := bson.M{"email": bson.M{"$eq": A.Email}}
	var dbUser User
	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("users")
	log.Printf("filter: %v", filter)
	err := collection.FindOne(*ctx, filter).Decode(&dbUser)
	if err != nil {
		return CreateResponse("failure", "error wrong email or password ", err.Error())
	}
	log.Printf("dbUser: %v", dbUser)
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(A.Password))
	if err != nil {
		return CreateResponse("failure", "bcrypt error ", err.Error())
	}
	var session Session
	res := session.Create(ctx, *A, dbUser)
	log.Printf("Session: %v", res)
	return CreateResponse("success", "success login", res)
}
