package handler

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	domain "userAuth-grpc/internal/domian"
	"userAuth-grpc/pkg/utils"
	"userAuth-grpc/proto/userAuth"
)

func (userAuthServer *UserAuthServer) GetUserAccountByPhone(ctx context.Context, in *userAuth.GetUserAccountByPhoneRequest) (*userAuth.GetUserAccountByPhoneResponse, error) {

	userAccount := userAuth.GetUserAccountByPhoneResponse{}
	userAccount.UserAccount = &userAuth.UserAccount{}
	userAccountDbm := &domain.UserAccount{}
	var err error
	if in.PhoneNumber != "" {
		err := userAccountDbm.GetUserAccountByPhone(ctx, in.PhoneNumber)
		if err == nil {
			err = utils.CopyFields(userAccountDbm, userAccount.UserAccount)
			if err != nil {
				return &userAccount, err
			}
			userAccount.UserAccount.Password = ""
		}
	} else {
		err = errors.New("phone number is empty")
	}

	return &userAccount, err
}

func (userAuthServer *UserAuthServer) VerifyUserEmailAndPassword(ctx context.Context, in *userAuth.VerifyUserEmailAndPasswordRequest) (*userAuth.GetUserAccountByPhoneResponse, error) {

	userAccount := userAuth.GetUserAccountByPhoneResponse{}
	userAccount.UserAccount = &userAuth.UserAccount{}
	userAccountDbm := &domain.UserAccount{}
	var err error
	if in.Email != "" {
		err := userAccountDbm.GetUserAccountByEmail(ctx, in.Email)
		if err == nil {
			err := bcrypt.CompareHashAndPassword([]byte(userAccountDbm.Password), []byte(in.Password))
			if err != nil {
				err = errors.New("password is incorrect")
				return &userAccount, err
			}
			err = utils.CopyFields(userAccountDbm, userAccount.UserAccount)
			if err != nil {
				return &userAccount, err
			}
			userAccount.UserAccount.Password = ""
		}
	} else {
		err = errors.New("email is empty")
	}

	return &userAccount, err
}

func (userAuthServer *UserAuthServer) CreateUserByEmail(ctx context.Context, in *userAuth.CreateUserByEmailRequest) (*userAuth.CreateUserByEmailResponse, error) {
	userAccount := userAuth.CreateUserByEmailResponse{}
	userAccount.UserAccount = &userAuth.UserAccount{}
	userAccountDbm := domain.UserAccount{}
	var err error
	if in.UserAccount.Email != "" {
		err := userAccountDbm.GetUserAccountByEmail(ctx, in.UserAccount.Email)
		if err == nil || userAccountDbm.UserID != 0 {
			err = errors.New("user already exists with this email")
		} else {
			hash, err := bcrypt.GenerateFromPassword([]byte(in.UserAccount.Password), bcrypt.DefaultCost)
			if err == nil {
				hashPassword := string(hash)
				err = utils.CopyFields(in.UserAccount, &userAccountDbm)
				if err == nil {
					userAccountDbm.Password = hashPassword
					userAccountDbm.Age = in.UserAccount.Age
					err = userAccountDbm.CreateUSerAccount(ctx)
					if err == nil {
						err = utils.CopyFields(&userAccountDbm, userAccount.UserAccount)
						if err == nil {
							userAccount.UserAccount.Password = ""
						}
					}
				}
			}
		}
	} else {
		err = errors.New("email is empty")
	}
	return &userAccount, err
}
