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
