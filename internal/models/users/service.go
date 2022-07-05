package users

import (
	"better-auth/internal/db"
	"better-auth/internal/models"
	"better-auth/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (U *User) Create(ctx *context.Context) map[string]interface{} {
	var isUser User
	filter := bson.M{"email": bson.M{"$eq": U.Email}}
	client := db.InitDB(*ctx)
	client.GetCollection("users").FindOne(*ctx, filter).Decode(&isUser)
	client.QuitDB()
	if (User{}) != isUser {
		return models.CreateResponse("failure", "user already registered", http.StatusConflict)
	}
	U.Base.GetBaseCreated()
	U.UserId = primitive.NewObjectID()
	U.ValidatorEmailSignIn = utils.StringToBase64(U.Email + U.Password)
	U.ValidatorUserNameSignIn = utils.StringToBase64(U.UserName + U.Password)
	U.Password = utils.GetHash([]byte(U.Password))
	dbClient := db.InitDB(*ctx)
	defer dbClient.DisconnectDB()
	res, err := dbClient.GetCollection("users").InsertOne(*ctx, U)
	if err != nil {
		return models.CreateResponse("Failure", "User creation failed", err)
	}
	return models.CreateResponse("Success", "User created successfully", map[string]interface{}{
		"_id":     res.InsertedID,
		"user_id": U.UserId,
	})
}

func (U *User) FindOne(ctx *context.Context, userID string) map[string]interface{} {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.CreateResponse("Failure", "User ID is incorrect", err)
	}
	filter := bson.M{"user_id": bson.M{"$eq": id}}
	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("users")
	cur := collection.FindOne(*ctx, filter)
	err = cur.Decode(&U)

	if err != nil {
		return models.CreateResponse("failure", "error finding user to database", err.Error())
	}
	return models.CreateResponse("success", "success finding user to database", U.ToMap())
}

func (U *User) Update(ctx *context.Context, userID string) map[string]interface{} {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.CreateResponse("failure", "error wrong user id", err.Error())
	}
	filter := bson.M{"user_id": bson.M{"$eq": id}}
	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("users")
	U.Base.GetBaseUpdated()
	update := U.GetUserAsBsonDocument()
	fields := bson.M{"$set": update}
	result, err := collection.UpdateOne(*ctx, filter, fields)
	if err != nil {
		return models.CreateResponse("failure", "error updating user to database", err.Error())
	}
	return models.CreateResponse("success", "success updating user to database", result)
}

func (U *User) Delete(ctx *context.Context, userID string) map[string]interface{} {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.CreateResponse("failure", "error wrong user id", err.Error())
	}
	filter := bson.M{"user_id": bson.M{"$eq": id}}
	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("users")
	g_res := U.FindOne(ctx, userID)
	if g_res["status"] == "failure" {
		return g_res
	}
	d_result, err := collection.DeleteOne(*ctx, filter)
	if err != nil {
		return models.CreateResponse("failure", "error deleting user to database", err.Error())
	}

	deleted_collection := client.GetCollection("users_deleted")
	i_result, errr := deleted_collection.InsertOne(*ctx, g_res)
	if errr != nil {
		return models.CreateResponse("failure", "error deleting user from database", err.Error())
	}

	return models.CreateResponse("success", "success deleting user to database", map[string]interface{}{
		"result": d_result,
		"hex":    i_result.InsertedID,
	})
}
