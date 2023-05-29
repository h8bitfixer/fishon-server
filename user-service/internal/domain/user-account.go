package domain

type UserAccount struct {
	UserID      uint32 `json:"userID"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	Age         int32  `json:"age"`
	Gender      int32  `json:"gender"`
}
