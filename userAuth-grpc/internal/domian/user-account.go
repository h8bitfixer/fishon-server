package domain

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"userAuth-grpc/pkg/db"
)

type UserAccount struct {
	UserID      uint32         `gorm:"column:user_id;primaryKey;autoIncrement"`
	Name        string         `gorm:"column:name;not null"`
	PhoneNumber string         `gorm:"column:phone_number;unique;not null"`
	Email       string         `gorm:"column:email;unique;not null"`
	Password    string         `gorm:"column:password;not null"`
	Age         int32          `gorm:"column:ago"`
	Gender      int32          `gorm:"column:gender"`
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

func (userAccount *UserAccount) GetUserAccountByPhone(_ context.Context, phoneNumber string) error {
	err := db.GetMySQLDB().Where("phone_number = ?", phoneNumber).First(userAccount).Error
	return err
}

func (userAccount *UserAccount) GetUserAccountByEmail(_ context.Context, email string) error {
	err := db.GetMySQLDB().Where("email = ?", email).First(userAccount).Error
	return err
}

func (userAccount *UserAccount) CreateUSerAccount(_ context.Context) error {
	err := db.GetMySQLDB().Create(userAccount).Error
	return err
}
