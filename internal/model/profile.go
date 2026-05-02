package model

import "time"

type Profile struct {
	ID         uint64    `json:"id"`
	UserID     uint64    `json:"user_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Phone      string    `json:"phone"`
	AvatarFileID *uint64 `json:"avatar_file_id,omitempty"`
	Gender     string    `json:"gender"`
	BirthDate  time.Time `json:"birth_date"`
	Bio        string    `json:"bio"`
	Country    string    `json:"country"`
	Province   string    `json:"province"`
	City       *string   `json:"city,omitempty"`
	District   *string   `json:"district,omitempty"`
	PostalCode *string   `json:"postal_code,omitempty"`
	Address    *string   `json:"address,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}