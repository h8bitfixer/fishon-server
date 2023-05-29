package domain

import (
	"context"
	"encoding/json"
	"time"
	"userAuth-grpc/pkg/db"
	//"userAuth-grpc/pkg/db"
)

type UserOTPRedisModel struct {
	UserPhoneNumber string `json:"userPhoneNumber"`
	UserOTP         string `json:"userOTP"`
	OTPGenerateTime int64  `json:"otpGenerateTime"`
}

// MarshalJSON serializes the UserOTPRedisModel struct to JSON and returns it as a string.
func (userOtpRM *UserOTPRedisModel) MarshalJSON() (string, error) {
	bytes, err := json.Marshal(userOtpRM)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// UnmarshalJSON deserializes the JSON data from a string into the UserOTPRedisModel struct.
func (userOtpRM *UserOTPRedisModel) UnmarshalJSON(data string) error {
	return json.Unmarshal([]byte(data), userOtpRM)
}

func (userOtpRM *UserOTPRedisModel) SetUserOTPRedisModel(ctx context.Context, pinToken string) error {
	userOtpRMStr, err := userOtpRM.MarshalJSON()
	if err == nil {
		err = db.GetRedisDB().Set(ctx, pinToken, userOtpRMStr, time.Minute*5).Err()
	}
	return err
}

func (userOtpRM *UserOTPRedisModel) GetUserOTPRedisModel(ctx context.Context, pinToken string) error {
	result, err := db.GetRedisDB().Get(ctx, pinToken).Result()
	if err == nil {
		err = userOtpRM.UnmarshalJSON(result)
	}
	return err
}
