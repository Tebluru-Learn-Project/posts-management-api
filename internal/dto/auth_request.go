package dto

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required|min=4"`
}

type RegisterRequest struct {
	Name	string `json:"name" binding:"required,min=3,max=225"`
	Email	string `json:"email" binding:"required,email"`
	Password	string `json:"password" binding:"required|min=6"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=6"`
}