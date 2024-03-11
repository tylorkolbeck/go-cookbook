package model

import "time"

type CookBook struct {
	Cookbook_id string    `json:"cookbook_id"`
	User_id     string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Public      bool      `json:"public"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
