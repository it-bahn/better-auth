package models

import (
	"better-auth/internal/db"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
CREATE USER
*/
func (U *User) Create(ctx *context.Context) map[string]interface{} {
	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("users")
	U.CreateNewUser()
	result, err := collection.InsertOne(*ctx, U)
	if err != nil {
		return CreateResponse("failure", "error inserting user to database", err.Error())
	}
	return CreateResponse("success", "success inserting user to database", map[string]interface{}{
		"user_id_from_result": result.InsertedID.(primitive.ObjectID).Hex(),
		"user_id":             U.UserId,
	})
}

/**
Update One User
*/
func (U *User) Update(ctx *context.Context, userID string) map[string]interface{} {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return CreateResponse("failure", "error wrong user id", err.Error())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}

	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("users")
	U.Base.GetBaseUpdated()

	update := U.GetUser()
	fields := bson.M{"$set": update}
	result, err := collection.UpdateOne(*ctx, filter, fields)
	if err != nil {
		return CreateResponse("failure", "error updating user to database", err.Error())
	}
	return CreateResponse("success", "success updating user to database", result)
}

/**
DELETE USER
*/
func (U *User) Delete(ctx *context.Context, userID string) map[string]interface{} {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return CreateResponse("failure", "error wrong user id", err.Error())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("users")
	result, err := collection.DeleteOne(*ctx, filter)
	if err != nil {
		return CreateResponse("failure", "error deleting user to database", err.Error())
	}
	return CreateResponse("success", "success deleting user to database", result)
}

/**
FIND ONE USER
*/
func (U *User) FindOne(ctx *context.Context, userID string) map[string]interface{} {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return CreateResponse("failure", "error wrong user id", err.Error())
	}
	filter := bson.M{"user_id": bson.M{"$eq": id}}
	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("users")
	cur := collection.FindOne(*ctx, filter)
	err = cur.Decode(&U)
	if err != nil {
		return CreateResponse("failure", "error finding user to database", err.Error())
	}
	return CreateResponse("success", "success finding user to database", U.ToMap())
}
