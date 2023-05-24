package domain

type OTPRequest struct {
	CountryCode string `json:"countryCode" binding:"required,max=4"`
	PhoneNumber string `json:"phoneNumber" binding:"required,max=11"`
	RequestID   string `json:"requestID"`
}
