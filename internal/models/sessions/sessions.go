package sessions

import (
	"better-auth/internal/models"
	"better-auth/internal/models/auth"
)

type Session struct {
	Base      models.BaseObject `json:"base,omitempty" bson:"base,omitempty"`
	SessionID string            `json:"session_id,omitempty" bson:"session_id,omitempty"`
	UserID    string            `json:"user_id,omitempty" bson:"user_id,omitempty"`
	AuthData  auth.AuthEmail    `json:"auth_data,omitempty" bson:"auth_data,omitempty"`
}

func (S *Session) IsEmptyEntirely() bool {
	return S.SessionID == "" && S.UserID == "" && S.AuthData.IsEmptyEntirely()
}
func (S *Session) IsEmptyMandatory() bool {
	return S.SessionID == "" || S.UserID == "" || S.AuthData.IsEmptyMandatory()
}
