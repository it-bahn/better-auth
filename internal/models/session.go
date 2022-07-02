package models

import (
	"better-auth/internal/db"
	"better-auth/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	Base      BaseObject `json:"base,omitempty" bson:"base,omitempty"`
	SessionID string     `json:"session_id,omitempty" bson:"session_id,omitempty"`
	UserID    string     `json:"user_id,omitempty" bson:"user_id,omitempty"`
	AuthData  AuthEmail  `json:"auth_data,omitempty" bson:"auth_data,omitempty"`
	JWTToken  string     `json:"jwt_token,omitempty" bson:"jwt_token,omitempty"`
}

func (S *Session) New(auth AuthEmail, user User) {
	S.Base.GetBaseCreated()
	S.AuthData.Email = auth.Email
	S.AuthData.Password = auth.Password
	S.SessionID = primitive.NewObjectID().Hex()
	S.UserID = user.UserId.Hex()
}
func (S *Session) Create(ctx *context.Context, auth AuthEmail, user User) map[string]interface{} {
	S.New(auth, user)
	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("sessions")
	S.AuthData.Password = utils.GetHash([]byte(S.AuthData.Password))
	res, err := collection.InsertOne(*ctx, S)
	if err != nil {
		return CreateResponse("failure", "error creating session", err.Error())
	}
	jwt, err := utils.GenerateJWT(user.ValidatorEmailSignIn)
	return CreateResponse("success", "success creating session", map[string]interface{}{
		"Inserted_ID": res.InsertedID,
		"Session_ID":  S.SessionID,
		"JWT_TOKEN":   jwt,
		"User_ID":     S.UserID,
		"User_Name":   "Welcome " + user.LastName + " , " + user.FirstName,
	})
}

func (S *Session) Delete(ctx *context.Context, SessionID string) map[string]interface{} {
	var dbSession Session
	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("sessions")
	filter := bson.M{"session_id": bson.M{"$eq": SessionID}}
	err := collection.FindOne(*ctx, filter).Decode(&dbSession)
	if err != nil {
		return CreateResponse("failure", "error deleting session", err.Error())
	}
	// Mo
	collection = client.GetCollection("sessions_deleted")
	dbSession.Base.GetBaseDeleted()
	_, err = collection.InsertOne(*ctx, dbSession)
	if err != nil {
		return CreateResponse("failure", "error deleting session", err.Error())
	}
	collection = client.GetCollection("sessions")
	res, err := collection.DeleteOne(*ctx, filter)
	return CreateResponse("success", "success deleting session", res)
}

func (S *Session) Get(ctx *context.Context, SessionID string) map[string]interface{} {
	client := db.InitDB(*ctx)
	defer client.DisconnectDB()
	collection := client.GetCollection("sessions")
	filter := bson.M{"session_id": bson.M{"$eq": SessionID}}
	var dbSession Session
	err := collection.FindOne(*ctx, filter).Decode(&dbSession)
	if err != nil {
		return CreateResponse("failure", "error getting session", err.Error())
	}
	return CreateResponse("success", "success getting session", dbSession)
}
