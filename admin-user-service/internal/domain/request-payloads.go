package domain

type OTPRequest struct {
	CountryCode string `json:"countryCode" binding:"required,max=4"`
	PhoneNumber string `json:"phoneNumber" binding:"required,max=11"`
	RequestID   string `json:"requestID"`
}

type OTPVerifyRequest struct {
	PinToken  string `json:"pinToken" binding:"required"`
	Otp       string `json:"otp" binding:"required"`
	RequestID string `json:"requestID"`
}

type CreateAccountRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int32  `json:"age"`
	Gender   int32  `json:"gender"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
