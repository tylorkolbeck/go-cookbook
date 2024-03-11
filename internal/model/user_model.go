package model

import "time"

type User struct {
	ID                string    `gorm:"primaryKey"`
	Email             string    `gorm:"uniqueIndex;not null"`
	Password          string    `gorm:"not null"` // Store hashed passwords only
	Created_at        time.Time `gorm:"autoCreateTime"`
	Updated_at        time.Time `gorm:"autoUpdateTime"`
	EmailVerified     bool      `json:"email_verified"`
	VerificationToken string    `json:"verification_token"`

	// Add other fields as necessary, e.g., Email, CreatedAt, etc.
}

type SafeUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	EmailVerified bool   `json:"email_verified"`
}
