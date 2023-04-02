package wd_models

import (
	"github.com/google/uuid"
)

type User struct {
	UserUUID uuid.UUID `json:"user_uuid,omitempty"`
	//HashedPassword string    `json:"hashed_password,omitempty"`
	Active       bool   `json:"active,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	//CreatedAt    time.Time `json:"created_at,omitempty"`
	//UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
