package users

import (
	"better-auth/internal/models"
	"better-auth/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type User struct {
	// Base
	Base                    models.BaseObject  `json:"base_object,omitempty" bson:"base_object,omitempty"`
	UserId                  primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	ValidatorEmailSignIn    string             `json:"validator_email,omitempty" bson:"validator_email,omitempty"`
	ValidatorUserNameSignIn string             `json:"validator_username,omitempty" bson:"validator_username,omitempty"`
	// '''''''''''''''''''''''''' REQUIRED FIELDS FOR USER ''''''''''''''''''''''''''
	FirstName string `json:"first_name,omitempty,required" bson:"first_name,omitempty,required"`
	LastName  string `json:"last_name,omitempty,required" bson:"last_name,omitempty,required"`
	UserName  string `json:"user_name,omitempty,required" bson:"user_name,omitempty,required"`
	Email     string `json:"email,omitempty,required" bson:"email,omitempty,required"`
	Password  string `json:"password,omitempty,required" bson:"password,omitempty,required"`
	// '''''''''''''''''''''''''' OPTIONAL FIELDS FOR USER ''''''''''''''''''''''''''
	Telephone   string             `json:"telephone,omitempty" bson:"telephone,omitempty"`
	Address     models.Address     `json:"address,omitempty" bson:"address,omitempty"`
	DateOfBirth models.DateOfBirth `json:"date_of_birth,omitempty" bson:"date_of_birth,omitempty"`
}

func (u *User) IsEmptyMandatory() bool {
	return u.FirstName == "" || u.LastName == "" || u.UserName == "" || u.Email == "" || u.Password == ""
}

func (u *User) IsEmptyEntirely() bool {
	return u.FirstName == "" && u.LastName == "" && u.UserName == "" && u.Email == "" && u.Password == "" && u.Telephone == "" && u.Address.IsEmptyEntirely() && u.DateOfBirth.IsEmptyEntirely()
}

func (U *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"user_id":       U.UserId,
		"first_name":    U.FirstName,
		"last_name":     U.LastName,
		"user_name":     U.UserName,
		"email":         U.Email,
		"telephone":     U.Telephone,
		"address":       U.Address,
		"date_of_birth": U.DateOfBirth,
	}
}

func (U *User) ToMapDetailed() map[string]interface{} {
	return map[string]interface{}{
		"base_object":        U.Base,
		"user_id":            U.UserId,
		"validator_email":    U.ValidatorEmailSignIn,
		"validator_username": U.ValidatorUserNameSignIn,
		"first_name":         U.FirstName,
		"last_name":          U.LastName,
		"user_name":          U.UserName,
		"email":              U.Email,
		"password":           U.Password,
		"telephone":          U.Telephone,
		"address":            U.Address,
		"date_of_birth":      U.DateOfBirth,
	}
}

func (U *User) GetUserAsBsonDocument() primitive.D {
	var res primitive.D
	res = append(res, primitive.E{Key: "base_object", Value: U.Base.GetAsBsonDocument()})
	if U.FirstName != "" {
		log.Printf("Updating FirstName.... %v", U.FirstName)
		res = append(res, primitive.E{Key: "first_name", Value: U.FirstName})
	}
	if U.Password != "" {
		log.Printf("Updating Password.... %v", U.Password)
		res = append(res, primitive.E{Key: "password", Value: utils.GetHash([]byte(U.Password))})
	}
	if U.LastName != "" {
		log.Printf("Updating LastName.... %v", U.LastName)
		res = append(res, primitive.E{Key: "last_name", Value: U.LastName})
	}
	if U.Email != "" {
		log.Printf("Updating Email.... %v", U.Email)
		res = append(res, primitive.E{Key: "email", Value: U.Email})
	}
	if U.Telephone != "" {
		log.Printf("Updating Phone.... %v", U.Telephone)
		res = append(res, primitive.E{Key: "phone", Value: U.Telephone})
	}
	if !U.DateOfBirth.IsEmptyEntirely() {
		log.Printf("Updating DateOfBirth.... %v", U.DateOfBirth)
		res = append(res, primitive.E{Key: "date_of_birth", Value: U.DateOfBirth.GetDOBAsBsonDocument()})
	}
	if !U.Address.IsEmptyEntirely() {
		log.Printf("Updating Address.... %v", U.Address)
		res = append(res, primitive.E{Key: "address", Value: U.Address.GetAddressAsBsonDocument()})
	}
	return res
}
