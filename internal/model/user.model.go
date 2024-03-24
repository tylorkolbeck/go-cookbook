package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"ID"`
	Email             string    `gorm:"not null" json:"email" gorm:"uniqueIndex;not null"`
	Password          string    `json:"password" gorm:"not null"` // Store hashed passwords only
	Created_at        time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	EmailVerified     bool      `json:"email_verified"`
	VerificationToken string    `json:"verification_token" gorm:"uniqueIndex;not null"`
	Name              string    `json:"name" gorm:"not null"`
	Role              string    `json:"role" gorm:"not null"`
	// Add other fields as necessary, e.g., Email, CreatedAt, etc.
}

func (u User) MarshalJSON() ([]byte, error) {
	type Alias User // Define an alias to avoid recursion
	return json.Marshal(&struct {
		*Alias
		Password          string `json:"-"` // Explicitly ignore the Password field
		VerificationToken string `json:"-"` // Explicitly ignore the VerificationToken field
	}{
		Alias: (*Alias)(&u),
	})
}

type SafeUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	EmailVerified bool   `json:"email_verified"`
	Role          string `json:"role"`
}
