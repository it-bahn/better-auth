package models

import (
	"better-auth/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type User struct {
	// Base
	Base                    BaseObject         `json:"base_object,omitempty" bson:"base_object,omitempty"`
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
	Telephone   string      `json:"telephone,omitempty" bson:"telephone,omitempty"`
	Address     Address     `json:"address,omitempty" bson:"address,omitempty"`
	DateOfBirth DateOfBirth `json:"date_of_birth,omitempty" bson:"date_of_birth,omitempty"`
}

func (u *User) IsEmptyMandatory() bool {
	return u.FirstName == "" || u.LastName == "" || u.UserName == "" || u.Email == "" || u.Password == ""
}

func (u *User) IsEmptyEntirely() bool {
	return u.FirstName == "" && u.LastName == "" && u.UserName == "" && u.Email == "" && u.Password == "" && u.Telephone == "" && u.Address.IsEmptyEntirely() && u.DateOfBirth.IsEmptyEntirely()
}

func (U *User) GetUser() primitive.D {
	var res primitive.D
	log.Printf("User: %v", U)
	if U.FirstName != "" {
		log.Printf("Updating FirstName.... %v", U.FirstName)
		res = append(res, primitive.E{Key: "first_name", Value: U.FirstName})
	}
	if U.LastName != "" {
		log.Printf("Updating LastName.... %v", U.LastName)
		res = append(res, primitive.E{Key: "last_name", Value: U.LastName})
	}
	if U.Email != "" {
		log.Printf("Updating Email.... %v", U.Email)
		res = append(res, primitive.E{Key: "email", Value: U.Email})
	}
	if U.Password != "" {
		log.Printf("Updating Password.... %v", U.Password)
		res = append(res, primitive.E{Key: "password", Value: utils.GetHash([]byte(U.Password))})
	}
	if U.Telephone != "" {
		log.Printf("Updating Telephone.... %v", U.Telephone)
		res = append(res, primitive.E{Key: "telephone", Value: U.Telephone})
	}
	if (DateOfBirth{}) != U.DateOfBirth {
		log.Printf("Updating DateOfBirth.... %v", U.DateOfBirth)
		res = append(res, primitive.E{Key: "date_of_birth", Value: U.DateOfBirth.GetDateOfBirth()})
	}
	if (Address{}) != U.Address {
		log.Printf("Updating Address.... %v", U.Address)
		res = append(res, primitive.E{Key: "address", Value: U.Address.GetAddress()})
	}
	return res
}
func (U *User) CreateNewUser() {
	U.Base.GetBaseCreated()
	U.UserId = primitive.NewObjectID()
	U.ValidatorEmailSignIn = utils.StringToBase64(U.Email + U.Password)
	U.ValidatorUserNameSignIn = utils.StringToBase64(U.UserName + U.Password)
	U.Password = utils.GetHash([]byte(U.Password))
}

func (U *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"user_id":       U.UserId,
		"first_name":    U.FirstName,
		"last_name":     U.LastName,
		"user_name":     U.UserName,
		"email":         U.Email,
		"telephone":     U.Telephone,
		"address":       U.Address.ToMap(),
		"date_of_birth": U.DateOfBirth.ToMap(),
	}
}
