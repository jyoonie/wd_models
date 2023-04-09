package wd_models

import (
	"github.com/google/uuid"
)

type Ingredient struct {
	IngredientUUID uuid.UUID `json:"ingredient_uuid,omitempty"`
	IngredientName string    `json:"ingredient_name,omitempty"`
	Category       string    `json:"category,omitempty"`
	DaysUntilExp   int       `json:"days_until_exp,omitempty"`
	//created_at           time.Time
	//updated_at           time.Time
}

type SearchIngredients struct {
	IngredientName string `json:"ingredient_name,omitempty"`
	Category       string `json:"category,omitempty"`
}
