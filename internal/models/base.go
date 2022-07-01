package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type BaseObject struct {
	BaseID    primitive.ObjectID `json:"base_id,omitempty" bson:"base_id,omitempty"`
	CreatedAt string             `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt string             `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Deleted   bool               `json:"deleted,omitempty" bson:"deleted,omitempty"`
}

type DateOfBirth struct {
	Day   int `json:"day,omitempty" bson:"day,omitempty"`
	Month int `json:"month,omitempty" bson:"month,omitempty"`
	Year  int `json:"year,omitempty" bson:"year,omitempty"`
}

type Address struct {
	Street  string `json:"street,omitempty" bson:"street,omitempty"`
	City    string `json:"city,omitempty" bson:"city,omitempty"`
	State   string `json:"state,omitempty" bson:"state,omitempty"`
	Zip     string `json:"zip,omitempty" bson:"zip,omitempty"`
	Country string `json:"country,omitempty" bson:"country,omitempty"`
}

func (A *Address) IsEmpty() bool {
	return A.Street == "" || A.City == "" || A.State == "" || A.Zip == "" || A.Country == ""
}
func (A *Address) IsEmptyEntirely() bool {
	return A.Street == "" && A.City == "" && A.State == "" && A.Zip == "" && A.Country == ""
}
func (D *DateOfBirth) IsEmpty() bool {
	return D.Day == 0 || D.Month == 0 || D.Year == 0
}
func (D *DateOfBirth) IsEmptyEntirely() bool {
	return D.Day == 0 && D.Month == 0 && D.Year == 0
}

func (B *BaseObject) IsEmpty() bool {
	return B.BaseID.Hex() == "" || B.CreatedAt == "" || B.UpdatedAt == ""
}
func (B *BaseObject) IsEmptyEntirely() bool {
	return B.BaseID.Hex() == "" && B.CreatedAt == "" && B.UpdatedAt == "" && B.Deleted == false
}

func (B *BaseObject) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"base_id":    B.BaseID,
		"created_at": B.CreatedAt,
		"updated_at": B.UpdatedAt,
	}
}

func (B *BaseObject) GetBaseCreated() {
	B.BaseID = primitive.NewObjectID()
	B.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	B.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	B.Deleted = false

}
func (B *BaseObject) GetBaseUpdated() {
	B.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
}
func (B *BaseObject) GetBaseDeleted() {
	B.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	B.Deleted = true
}

func (D *DateOfBirth) GetDateOfBirth() primitive.D {
	var res primitive.D
	log.Printf("DateOfBirth: %v", D)
	if D.Day != 0 {
		log.Printf("Updating Day.... %v", D.Day)
		res = append(res, primitive.E{Key: "day", Value: D.Day})
	}
	if D.Month != 0 {
		log.Printf("Updating Month.... %v", D.Month)
		res = append(res, primitive.E{Key: "month", Value: D.Month})

	}
	if D.Year != 0 {
		log.Printf("Updating Year.... %v", D.Year)
		res = append(res, primitive.E{Key: "year", Value: D.Year})
	}
	return res
}

func (D *DateOfBirth) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"day":   D.Day,
		"month": D.Month,
		"year":  D.Year,
	}
}

func (A *Address) GetAddress() primitive.D {
	var res primitive.D
	log.Printf("Address: %v", A)
	if A.Street != "" {
		log.Printf("Updating Street.... %v", A.Street)
		res = append(res, primitive.E{Key: "street", Value: A.Street})
	}
	if A.City != "" {
		log.Printf("Updating City.... %v", A.City)
		res = append(res, primitive.E{Key: "city", Value: A.City})
	}
	if A.State != "" {
		log.Printf("Updating State.... %v", A.State)
		res = append(res, primitive.E{Key: "state", Value: A.State})
	}
	if A.Zip != "" {
		log.Printf("Updating Zip.... %v", A.Zip)
		res = append(res, primitive.E{Key: "zip", Value: A.Zip})
	}
	if A.Country != "" {
		log.Printf("Updating Country.... %v", A.Country)
		res = append(res, primitive.E{Key: "country", Value: A.Country})
	}
	return res
}

func (A *Address) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"street":  A.Street,
		"city":    A.City,
		"state":   A.State,
		"zip":     A.Zip,
		"country": A.Country,
	}
}
