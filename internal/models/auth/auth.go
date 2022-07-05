package auth

import (
	"better-auth/internal/db"
	"better-auth/internal/models"
	"better-auth/internal/models/users"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthEmail struct {
	Email    string `json:"email,required" bson:"email,required"`
	Password string `json:"password,required" bson:"password,required"`
}

func (a *AuthEmail) IsEmptyMandatory() bool {
	return a.Email == "" || a.Password == ""
}

func (a *AuthEmail) IsEmptyEntirely() bool {
	return a.Email == "" && a.Password == ""
}

func (A *AuthEmail) Login(ctx *context.Context) map[string]interface{} {
	filter := bson.M{"email": bson.M{"$eq": A.Email}}
	var dbUser users.User
	client := db.InitDB(*ctx)

	collection := client.GetCollection("users")
	log.Printf("filter: %v", filter)
	err := collection.FindOne(*ctx, filter).Decode(&dbUser)
	if err != nil {
		return models.CreateResponse("failure", "error email not registered!", err.Error())
	}
	log.Printf("dbUser: %v", dbUser)
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(A.Password))
	if err != nil {
		return models.CreateResponse("failure", "error incorrect email or password", err.Error())
	}
	client.QuitDB()
	return map[string]interface{}{
		"status":  "success",
		"user_id": dbUser.UserId.Hex(),
	}
}
