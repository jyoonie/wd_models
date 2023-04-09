package wd_models

import (
	"time"

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

type FridgeIngredient struct {
	UserUUID       uuid.UUID `json:"user_uuid,omitempty"`
	IngredientUUID uuid.UUID `json:"ingredient_uuid,omitempty"`
	Amount         int       `json:"amount,omitempty"`
	Unit           string    `json:"unit,omitempty"`
	PurchasedDate  time.Time `json:"purchased_date,omitempty"`
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	// created_at       time.Time
	// updated_at       time.Time
}

type DeleteFIngr struct {
	UserUUID       uuid.UUID `json:"user_uuid,omitempty"`
	IngredientUUID uuid.UUID `json:"ingredient_uuid,omitempty"`
}
