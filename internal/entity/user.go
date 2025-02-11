package entity

import (
	"time"
)

type User struct {
	ID        uint         `json:"id" gorm:"primaryKey"`
	Name      string       `json:"name"`
	Email     string       `json:"email" gorm:"unique;not null"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Auth      ExternalAuth `gorm:"foreignKey:UserID"`
}

type ExternalAuth struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	UserID         uint      `json:"user_id" gorm:"unique;not null"`
	Provider       string    `json:"provider" gorm:"not null"`
	ProviderUserID string    `json:"provider_user_id" gorm:"unique;not null"`
	AccessToken    string    `json:"access_token"`
	RefreshToken   string    `json:"refresh_token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
