package dto

type UpdateProfileRequest struct {
	AvatarFileID *uint64 `json:"avatar_file_id"`
	FirstName  string  `json:"first_name" binding:"required,min=2,max=100"`
	LastName   string  `json:"last_name" binding:"required,min=2,max=100"`
	Phone      string  `json:"phone" binding:"required,min=8,max=20"`
	Gender     string  `json:"gender" binding:"required,oneof=male female"`
	BirthDate  string  `json:"birth_date" binding:"required,datetime=2006-01-02"`
	Bio        string  `json:"bio" binding:"required"`
	Country    string  `json:"country" binding:"required"`
	Province   string  `json:"province" binding:"required"`
	City       *string `json:"city"`
	District   *string `json:"district"`
	PostalCode *string `json:"postal_code"`
	Address    *string `json:"address"`
}