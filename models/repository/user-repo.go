package repository

import (
	"time"
)

type User struct {
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Address   string     `json:"address,omitempty"`
	Phone     string     `json:"phone,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
