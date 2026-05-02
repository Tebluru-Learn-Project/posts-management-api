package dto

type ProfileResponse struct {
	ID           uint64  `json:"id"`
	UserID       uint64  `json:"user_id"`
	AvatarFileID *uint64 `json:"avatar_file_id,omitempty"`
	AvatarURL    *string `json:"avatar_url,omitempty"`

	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Phone      string  `json:"phone"`
	Gender     string  `json:"gender"`
	BirthDate  string  `json:"birth_date"`
	Bio        string  `json:"bio"`
	Country    string  `json:"country"`
	Province   string  `json:"province"`
	City       *string `json:"city,omitempty"`
	District   *string `json:"district,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	Address    *string `json:"address,omitempty"`
}