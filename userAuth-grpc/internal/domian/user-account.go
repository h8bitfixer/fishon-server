package domain

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"userAuth-grpc/pkg/db"
)

type UserAccount struct {
	UserID      uint           `gorm:"column:user_id;primaryKey;autoIncrement"`
	Name        string         `gorm:"column:name;not null"`
	PhoneNumber string         `gorm:"column:phone_number;unique;not null"`
	Email       string         `gorm:"column:email;unique;not null"`
	Password    string         `gorm:"column:password;not null"`
	Age         int64          `gorm:"column:ago"`
	Gender      int64          `gorm:"column:gender"`
	CreatedAt   int64          `gorm:"column:created_at"`
	UpdatedAt   int64          `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

// MarshalJSON serializes the struct to JSON and returns it as a string.
func (userAccount *UserAccount) MarshalJSON() (string, error) {
	bytes, err := json.Marshal(userAccount)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// UnmarshalJSON deserializes the JSON data from a string
func (userAccount *UserAccount) UnmarshalJSON(data string) error {
	return json.Unmarshal([]byte(data), userAccount)
}

func (userAccount *UserAccount) GetUserAccountByPhone(ctx context.Context, phoneNumber string) error {
	err := db.GetMySQLDB().Where("phone = ?", phoneNumber).First(userAccount).Error
	return err
}
