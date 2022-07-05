package sessions

import (
	"better-auth/internal/db"
	"better-auth/internal/models"
	"better-auth/internal/models/auth"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (S *Session) Create(ctx *context.Context, Email string, userId string, auth auth.AuthEmail) map[string]interface{} {
	dbClient := db.InitDB(*ctx)
	defer dbClient.DisconnectDB()
	S.SessionID = primitive.NewObjectID().Hex()
	S.UserID = userId
	S.Base.GetBaseCreated()
	S.AuthData.Email = Email
	res, err := dbClient.GetCollection("sessions").InsertOne(*ctx, S)
	if err != nil {
		return models.CreateResponse("failure", "error creating session", err.Error())
	}
	return map[string]interface{}{
		"_id":        res.InsertedID,
		"session_id": S.SessionID,
		"user_id":    S.UserID,
		"timestamp":  time.Now().Format("2006-01-02 15:04:05"),
	}
}
func (S *Session) FindOne(ctx *context.Context, session_id string) map[string]interface{} {
	var dbSession Session
	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("sessions")
	filter := bson.M{"session_id": bson.M{"$eq": session_id}}
	err := collection.FindOne(*ctx, filter).Decode(&dbSession)
	if err != nil {
		return models.CreateResponse("failure", "error finding session", err.Error())
	}
	return models.CreateResponse("success", "Session Found!", map[string]interface{}{
		"session_id": dbSession.SessionID,
		"user_id":    dbSession.UserID,
	})
}

func (S *Session) Delete(ctx *context.Context, session_id string) map[string]interface{} {
	var dbSession Session
	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("sessions")
	filter := bson.M{"session_id": bson.M{"$eq": session_id}}
	err := collection.FindOne(*ctx, filter).Decode(&dbSession)
	if err != nil {
		return models.CreateResponse("failure", "error deleting session", err.Error())
	}
	collection = client.GetCollection("sessions_deleted")
	dbSession.Base.GetBaseDeleted()
	_, err = collection.InsertOne(*ctx, dbSession)
	if err != nil {
		return models.CreateResponse("failure", "error deleting session", err.Error())
	}
	collection = client.GetCollection("sessions")
	res, err := collection.DeleteOne(*ctx, filter)
	if err != nil {
		return models.CreateResponse("failure", "error deleting session", err.Error())
	}
	return models.CreateResponse("success", "Logged Out Successfully", res.DeletedCount)
}
